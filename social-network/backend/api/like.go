package api

import (
	"social-network/cache"
	"social-network/middleware/logger"
	"social-network/models"
	"social-network/queue"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LikeAPI struct{}

type LikeRequest struct {
	UserID   uint64 `json:"user_id"`   // 用户ID
	MomentID uint64 `json:"moment_id"` // 朋友圈消息ID
}

func (l LikeAPI) LikeMomentBySQL(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	userid := req.UserID
	momentid := req.MomentID

	id, err := models.AddLike(userid, momentid)
	if err != nil {
		ReturnError(c, 4002, "添加失败")
		return
	}
	ReturnSucess(c, 2000, "点赞成功", id, 1)

}

func (l LikeAPI) LikeMomentByRedis(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	userID := strconv.FormatUint(req.UserID, 10)
	momentID := strconv.FormatUint(req.MomentID, 10)

	liked, err := cache.Rdb.SIsMember(cache.Rctx, "likes:"+momentID, userID).Result()
	if err != nil {
		ReturnError(c, 4002, "点赞失败")
		return
	}
	if liked {
		ReturnSucess(c, 2001, "您已经点赞了", momentID, 1)
		return
	}

	_, err = cache.Rdb.SAdd(cache.Rctx, "likes:"+momentID, userID).Result()
	if err != nil {
		ReturnError(c, 4002, "点赞失败")
		return
	}
	_, err = cache.Rdb.HIncrBy(cache.Rctx, "likes:count", momentID, 1).Result()
	if err != nil {
		ReturnError(c, 4003, "点赞失败")
		return
	}
	if err = queue.PublishLikeAction(req.UserID, req.MomentID, true); err != nil {
		logger.Error(map[string]interface{}{"RabbitMQ Publisg error": err.Error()})
		ReturnError(c, 4004, "点赞失败")
	}
	ReturnSucess(c, 2000, "点赞成功", req, 1)

}

func (l LikeAPI) UnLikeMomentByRedis(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}

	userID := strconv.FormatUint(req.UserID, 10)
	momentID := strconv.FormatUint(req.MomentID, 10)

	liked, err := cache.Rdb.SIsMember(cache.Rctx, "likes:"+momentID, userID).Result()
	if err != nil {
		ReturnError(c, 4002, "取消失败")
		return
	}
	if !liked {
		ReturnSucess(c, 2001, "您未点赞", momentID, 1)
		return
	}
	_, err = cache.Rdb.SRem(cache.Rctx, "likes:"+momentID, userID).Result()
	if err != nil {
		ReturnError(c, 4002, "取消失败")
		return
	}
	_, err = cache.Rdb.HIncrBy(cache.Rctx, "likes:count", momentID, -1).Result()
	if err != nil {
		ReturnError(c, 4003, "取消失败")
		return
	}
	err = queue.PublishLikeAction(req.UserID, req.MomentID, false)
	if err != nil {
		ReturnError(c, 4004, "取消失败")
		return
	}

	ReturnSucess(c, 2000, "取消成功", req, 1)

}

func (l LikeAPI) GetLikesByMoment(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	momentID := strconv.FormatUint(req.MomentID, 10)
	liks, err := cache.Rdb.HGet(cache.Rctx, "likes:count", momentID).Int64()
	if err == redis.Nil {
		liks, err := models.GetLikes(req.MomentID)
		if err != nil {
			ReturnError(c, 4002, "查询失败")
			return
		}

		count := uint64(len(liks))
		_, err = cache.Rdb.HSet(cache.Rctx, "likes:count", momentID, count).Result()
		if err != nil {
			queue.PublishLikeCounts(req.MomentID, count)
		}
		ReturnSucess(c, 2000, "查询成功", count, 1)
		return
	}
	ReturnSucess(c, 2001, "查询成功", liks, 1)
}
