package api

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime"
	"net/url"
	"os"
	"path/filepath"
	"social-network/cache"
	"social-network/middleware/logger"
	"social-network/models"
	"social-network/queue"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MomentAPI struct{}

// 请求对象
type MomentReq struct {
	MomentID   uint64 `json:"id"`
	UserID     uint64 `json:"user_id"`
	Content    string `json:"content"`
	ImagePaths string `json:"image_paths"`
}

func (m MomentAPI) SendMoment(c *gin.Context) {
	var req MomentReq
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "提交内容有误")
		return
	}
	userid := req.UserID
	content := req.Content
	img := req.ImagePaths
	if userid == 0 || content == "" {
		ReturnError(c, 4001, "提交内容有误")
		return
	}

	user, err := models.GetUserInfoByID(userid)
	if err != nil {
		ReturnError(c, 4002, "提交失败")
		return
	}
	if user.Id == 0 {
		ReturnError(c, 4001, "用户不存在")
		return
	}
	err = models.AddMoment(userid, content, img)
	if err != nil {
		ReturnError(c, 4002, "提交失败")
		return
	}
	ReturnSucess(c, 0, "提交成功", req, 1)
}

func generateRandomFilename(extension string) (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x%s", b, extension), nil
}

func (api MomentAPI) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		ReturnError(c, 4001, "上传失败")
		return
	}
	defer file.Close()
	saveDir := "./uploads/"
	err = os.MkdirAll(saveDir, 0755)
	if err != nil {
		ReturnError(c, 4002, "创建目录失败: "+err.Error())
		return
	}

	// 提取文件扩展名
	ext := filepath.Ext(header.Filename)
	if ext == "" { // 如果没有扩展名，尝试从MIME类型推断
		mimeType := mime.TypeByExtension(header.Filename)
		exts, _ := mime.ExtensionsByType(mimeType)
		if len(exts) > 0 {
			ext = exts[0]
		}
	}

	// 生成随机文件名
	randomFilename, err := generateRandomFilename(ext)
	if err != nil {
		ReturnError(c, 4002, "生成文件名失败: "+err.Error())
		return
	}
	savePath := filepath.Join(saveDir, randomFilename)
	out, err := os.Create(savePath)
	if err != nil {
		ReturnError(c, 4002, "文件创建失败: "+err.Error())
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		ReturnError(c, 4003, "文件保存失败: "+err.Error())
		return
	}
	data := filepath.Join("/uploads", randomFilename)
	ReturnSucess(c, 0, "上传成功", data, 1)
}

// 返回对象
type MomInfo struct {
	models.Moment
	Username   string `json:"username"`
	Likestatus bool   `json:"like_status"`
	Likescount uint64 `json:"likes_count"`
}

func (m MomentAPI) GetPageMomentList(c *gin.Context) {
	var req PageReq
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	userid := req.UserID

	moments, err := models.GetPageMomentByTime(req.Counts, req.Offset)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	var Results []MomInfo
	for _, moment := range moments {
		//根据内容表匹配 作者
		author, err := models.GetUserInfoByID(moment.UserID)
		if err != nil {
			ReturnError(c, 4003, "查询失败")
			return
		}
		authorName := author.Username
		//确认like_status的状态 有就是有，没有就是没有点赞，这个数据一定是同步的
		//在新redis节点上 可能出现异常
		strMomID := strconv.FormatUint(moment.ID, 10)
		keySRedis := "likes:" + strMomID
		like_status, err := models.CheckLikeStatus(moment.ID, userid)
		if err != nil {
			ReturnError(c, 4004, "查询失败")
			return
		}
		//根据内容匹配 确定点赞数
		//先查 redis 该文章的点赞记录  包含likes:MomID键 点赞记录 = SET(likes:MomID) 的成员数
		likes_count, err := cache.Rdb.SCard(cache.Rctx, keySRedis).Result()
		if err != nil {
			ReturnError(c, 4005, "查询失败")
			return
		}

		//redis中没有该moment的数据   去到数据库确定
		if likes_count == 0 {
			likes, err := models.GetLikes(moment.ID)
			if err != nil {
				ReturnError(c, 4006, "查询失败")
				return
			}
			likes_count = int64(len(likes))
			//同步到更新 redis 单条点赞记录 增加到redis
			go func() {

				//删除和该moment有关的所有键
				err := cache.DelMonent(moment.ID)
				if err != nil {
					//删除失败 mq处理
					msg := queue.DelMsg{
						MomentID: moment.ID,
					}
					err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
					if err != nil {
						logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
					}
				}
				err = cache.AddLikes(likes)
				if err != nil {
					msg := queue.DelMsg{
						MomentID: moment.ID,
					}
					err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
					if err != nil {
						logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
					}
				}

			}()
		}
		Results = append(Results, MomInfo{
			Moment:     moment,
			Username:   authorName,
			Likescount: uint64(likes_count),
			Likestatus: like_status,
		})

	}
	ReturnSucess(c, 0, "查询成功", Results, int64(len(Results)))
}

func (m MomentAPI) GetMomentsByUserID(c *gin.Context) {
	var req PageReq
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	userid := uint64(req.UserID)
	if userid == 0 {
		ReturnError(c, 4001, "参数错误")
		return
	}

	if userid < 5 {
		//管理员账号
		moments, err := models.GetPageMomentByTime(req.Counts, req.Offset)
		if err != nil {
			ReturnError(c, 4002, "查询失败")
			return
		}
		ReturnSucess(c, 0, "查询成功", moments, int64(len(moments)))
		return
	}

	moments, err := models.GetUserMoments(userid, req.Counts, req.Offset)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}

	var Results []MomInfo
	author, err := models.GetUserInfoByID(userid)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	authorName := author.Username
	for _, moment := range moments {
		//确认like_status的状态
		strMomID := strconv.FormatUint(moment.ID, 10)
		keySRedis := "likes:" + strMomID
		like_status, err := models.CheckLikeStatus(moment.ID, userid)
		if err != nil {
			ReturnError(c, 4002, "查询失败")
			return
		}
		//根据内容匹配 确定点赞数
		//先查 redis 该文章的点赞记录  包含likes:MomID键 点赞记录 = SET(likes:MomID) 的成员数
		likes_count, err := cache.Rdb.SCard(cache.Rctx, keySRedis).Result()
		if err != nil {
			ReturnError(c, 4002, "查询失败")
			return
		}
		//redis中没有  去到数据库中确认
		if likes_count == 0 {
			likes, err := models.GetLikes(moment.ID)
			if err != nil {
				ReturnError(c, 4002, "查询失败")
				return
			}
			likes_count = int64(len(likes))
			//同步到更新 redis 单条点赞记录 增加到redis
			go func() {
				//删除和该moment有关的所有键
				err := cache.DelMonent(moment.ID)
				if err != nil {
					//删除失败 mq处理
					msg := queue.DelMsg{
						MomentID: moment.ID,
					}
					err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
					if err != nil {
						logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
					}
				}
				err = cache.AddLikes(likes)
				if err != nil {
					msg := queue.DelMsg{
						MomentID: moment.ID,
					}
					err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
					if err != nil {
						logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
					}
				}
			}()
		}
		Results = append(Results, MomInfo{
			Moment:     moment,
			Username:   authorName,
			Likescount: uint64(likes_count),
			Likestatus: like_status,
		})

	}
	ReturnSucess(c, 0, "查询成功", Results, int64(len(Results)))
}

func (m MomentAPI) DeleteMoment(c *gin.Context) {
	var req MomentReq
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	momentID := req.MomentID
	if momentID == 0 {
		ReturnError(c, 4001, "参数有误")
	}
	//获取moment图片对应的url

	moment, err := models.GetMomentByID(momentID)
	if err != nil {
		ReturnError(c, 4002, "删除失败")
		return
	}
	if err := models.DeleteMoment(momentID); err != nil {
		ReturnError(c, 4004, "删除失败")
		return
	}
	ReturnSucess(c, 0, "删除成功", momentID, 1)
	//异步删除图片
	go func() {
		//删除服务器中对应的内容图片
		if moment.ImagePaths != "" {
			paths := strings.Split(moment.ImagePaths, ",")
			for _, path := range paths {
				if path != "" {
					// 删除图片文件
					parsedUrl, err := url.Parse(path)
					if err != nil {
						logger.Error(map[string]interface{}{"parsedUrl  err:": err})
						continue
					}
					filepath := "." + parsedUrl.Path
					err = os.Remove(filepath)
					if err != nil {
						logger.Error(map[string]interface{}{"Failed to delete image file:": err})
					}
				}
			}
		}

	}()
	//异步更新redis
	go func() {
		//删除和该moment有关的所有键
		err := cache.DelMonent(momentID)
		if err != nil {
			//删除失败 mq处理
			msg := queue.DelMsg{
				MomentID: momentID,
			}
			err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
			if err != nil {
				logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
			}
		}
	}()

}

func (m MomentAPI) RankMoments(c *gin.Context) {
	var req PageReq
	var Results []MomInfo
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数出错")
		return
	}
	userid := req.UserID

	exists, err := cache.Rdb.Exists(cache.Rctx, "likes:ranked").Result()
	if err != nil {
		ReturnError(c, 4002, "查询出错")
		return
	}
	if exists == 1 { //有就直接装载数据
		size, err := cache.Rdb.ZCard(cache.Rctx, "likes:ranked").Result()
		if err != nil {
			ReturnError(c, 4002, "查询出错")
			return
		}
		start := req.Offset
		end := req.Offset + req.Counts
		if int64(end) >= size {
			end = int(size)
		}
		resRedis, err := cache.Rdb.ZRevRangeWithScores(cache.Rctx, "likes:ranked", int64(start), int64(end)).Result()
		if err != nil {
			ReturnError(c, 4002, "查询出错")
			return
		}
		for _, zRedis := range resRedis {
			momentid, err := strconv.ParseUint(zRedis.Member.(string), 10, 64)
			if err != nil {
				ReturnError(c, 4002, "查询出错")
				return
			}
			moment, err := models.GetMomentByID(momentid)
			if err != nil {
				ReturnError(c, 4002, "查询出错")
				return
			}
			user, err := models.GetUserInfoByID(moment.UserID)
			if err != nil {
				ReturnError(c, 4002, "查询出错")
				return
			}
			username := user.Username

			//确认like_status的状态
			like_status, err := models.CheckLikeStatus(moment.ID, userid)
			if err != nil {
				ReturnError(c, 4002, "查询失败")
				return
			}

			Results = append(Results, MomInfo{
				Moment:     moment,
				Username:   username,
				Likestatus: like_status,
				Likescount: uint64(zRedis.Score),
			})
		}
		ReturnSucess(c, 0, "查询成功", Results, int64(len(Results)))
		return
	} else {
		//没有就查询数据库
		likeinfos, err := models.RankMoments(req.Counts, req.Offset)
		if err != nil {
			ReturnError(c, 4002, "查询出错")
			return
		}
		//匹配其他信息 作者名字 点赞状态
		for _, likeinfo := range likeinfos {
			moment, err := models.GetMomentByID(likeinfo.MomentID)
			if err != nil {
				ReturnError(c, 4002, "查询出错")
				return
			}
			likes_count := likeinfo.Likescount
			author, err := models.GetUserInfoByID(moment.UserID)
			if err != nil {
				ReturnError(c, 4002, "查询出错")
				return
			}
			authorName := author.Username
			like_status, err := models.CheckLikeStatus(moment.ID, moment.UserID)
			if err != nil {
				ReturnError(c, 4002, "查询失败")
				return
			}

			Results = append(Results, MomInfo{
				Moment:     moment,
				Username:   authorName,
				Likestatus: like_status,
				Likescount: likes_count,
			})

		}
		ReturnSucess(c, 0, "查询成功", Results, int64(len(Results)))

		go func() {
			//删除和该moment有关的所有键
			for _, likeinfo := range likeinfos {
				momentID := likeinfo.MomentID
				err := cache.DelMonent(momentID)
				if err != nil {
					//删除失败 mq处理
					msg := queue.DelMsg{
						MomentID: momentID,
					}
					err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
					if err != nil {
						logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
					}
				}
				//删除完成，加入点赞信息
				likes, err := models.GetLikes(momentID)
				if err != nil {
					return
				}
				err = cache.AddLikes(likes)
				if err != nil {
					//事务加入出错，用队列进行删除，确保数据一致性
					msg := queue.DelMsg{
						MomentID: momentID,
					}
					err := queue.PublishMsg(msg, queue.QDelKey_ToRedis)
					if err != nil {
						logger.Error(map[string]interface{}{"RabbitMQ CHDelKey_ToRedis err:": err})
					}
				}
			}

		}()

	}
}
