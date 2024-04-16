package api

import (
	"social-network/cache"
	"social-network/middleware/logger"
	"social-network/models"
	"social-network/queue"

	"github.com/gin-gonic/gin"
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
	if req.MomentID == 0 || req.UserID == 0 {
		ReturnError(c, 4001, "参数有误")
		return
	}
	id, err := models.AddLike(userid, momentid)
	if err != nil {
		ReturnError(c, 4002, "添加失败")
		return
	}
	ReturnSucess(c, 0, "点赞成功", id, 1)
}

func (l LikeAPI) UnLikeMomentBySQL(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	userid := req.UserID
	momentid := req.MomentID
	if req.MomentID == 0 || req.UserID == 0 {
		ReturnError(c, 4001, "参数有误")
		return
	}
	err := models.DelLike(userid, momentid)
	if err != nil {
		ReturnError(c, 4002, "删除失败")
		return
	}
	ReturnSucess(c, 0, "删除成功", momentid, 1)
}

func (l LikeAPI) LikeMomentByRedis(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}

	if req.MomentID == 0 || req.UserID == 0 {
		ReturnError(c, 4001, "参数有误")
		return
	}

	likes_count, err := cache.AddLike(req.MomentID, req.UserID)

	if err != nil {
		ReturnError(c, 4003, "点赞失败")
		return
	}

	if likes_count != -1 {
		ReturnSucess(c, 0, "点赞成功", likes_count, 1)
		//异步写 mysql
		go func() {
			_, err := models.AddLike(req.UserID, req.MomentID)
			if err != nil {
				//写mysq失败 改用 消息队列
				msg := queue.LikeMsg{
					UserID:   req.UserID,
					MomentID: req.MomentID,
					Action:   true,
				}
				err = queue.PublishMsg(msg, queue.QAction_ToSQL)
				if err != nil {
					logger.Error(map[string]interface{}{"QAction_ToSQL error :": err.Error()})
				}
			}
		}()
		return
	}
	ReturnError(c, 4003, "点赞失败")
}

func (l LikeAPI) UnLikeMomentByRedis(c *gin.Context) {
	var req LikeRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	if req.MomentID == 0 || req.UserID == 0 {
		ReturnError(c, 4001, "参数有误")
		return
	}

	likes_count, err := cache.DelLike(req.MomentID, req.UserID)
	if err != nil {
		ReturnError(c, 4002, "删除失败")
		return
	}

	if likes_count != -1 {
		ReturnSucess(c, 0, "取消成功", likes_count, 1)

		go func() {
			err := models.DelLike(req.UserID, req.MomentID)
			if err != nil {
				//写mysql失败 改用 消息队列
				msg := queue.LikeMsg{
					UserID:   req.UserID,
					MomentID: req.MomentID,
					Action:   false,
				}
				err = queue.PublishMsg(msg, queue.QAction_ToSQL)
				if err != nil {
					logger.Error(map[string]interface{}{"UnLikeActionToSQL error :": err.Error()})
					return
				}
			}
		}()
		return
	}
	ReturnError(c, 4004, "删除失败")

}
