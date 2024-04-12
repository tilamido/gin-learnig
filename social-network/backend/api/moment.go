package api

import (
	"social-network/models"

	"github.com/gin-gonic/gin"
)

type MomentAPI struct{}

type SendRequest struct {
	UserID     uint64 `gorm:"index;notNull;column:user_id" json:"user_id"`
	Content    string `gorm:"type:text;notNull;column:content" json:"content"`
	ImagePaths string `gorm:"type:text;column:image_paths" json:"image_paths"`
}

func (m MomentAPI) SendMoment(c *gin.Context) {
	var sendreq SendRequest
	if err := c.BindJSON(&sendreq); err != nil {
		ReturnError(c, 4001, "提交内容有误")
		return
	}

	id, err := models.AddMoment(sendreq.UserID, sendreq.Content, sendreq.ImagePaths)
	if err != nil {
		ReturnError(c, 4002, "提交失败")
		return
	}

	ReturnSucess(c, 2000, "提交成功", id, 1)
}

type MomentListRequest struct {
	Counts int `json:"counts"`
	Offset int `json:"offset"`
}

func (m MomentAPI) GetPageMomentList(c *gin.Context) {
	var req MomentListRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}

	moments, err := models.GetPageMomentListByTime(req.Counts, req.Offset)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	counts := len(moments)
	ReturnSucess(c, 2000, "读取成功", moments, int64(counts))

}

func (m MomentAPI) GetAllMomentList(c *gin.Context) {
	moments, err := models.GetAllMomentListByTime()
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	counts := len(moments)
	ReturnSucess(c, 2000, "读取成功", moments, int64(counts))

}

type MomentsByUserIDRequest struct {
	UserID   int `json:"user_id"`
	MomentId int `json:"moment_id"`
}

func (m MomentAPI) GetMomentsByUserID(c *gin.Context) {
	var req MomentsByUserIDRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	userid := uint64(req.UserID)
	moments, err := models.GetUserMoments(userid)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}

	counts := len(moments)
	ReturnSucess(c, 2000, "读取成功", moments, int64(counts))
}

func (m MomentAPI) DeleteMoment(c *gin.Context) {

	var req MomentsByUserIDRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	momentID := uint64(req.MomentId)

	if err := models.DeleteMoment(momentID); err != nil {
		ReturnError(c, 4002, "数据库删除失败")
	}

}
