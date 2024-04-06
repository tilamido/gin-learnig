package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidstr := c.DefaultPostForm("aid", "-1")
	aid, err := strconv.Atoi(aidstr)
	if aid == -1 || err != nil {
		ReturnError(c, 0, "活动不存在")
		return
	}

	players, err := models.GetPlayersByAID(aid)
	if err != nil {
		ReturnError(c, 401, "查询失败")
		return
	}

	ReturnSucess(c, 0, "查询成功", players, 1)

}

func (p PlayerController) GetPlayersRank(c *gin.Context) {
	aidstr := c.DefaultPostForm("aid", "-1")
	aid, err := strconv.Atoi(aidstr)
	if aid == -1 || err != nil {
		ReturnError(c, 0, "活动不存在")
		return
	}

	players, err := models.GetPlayersRankByAID(aid, "score desc")
	if err != nil {
		ReturnError(c, 401, "查询失败")
		return
	}

	ReturnSucess(c, 0, "查询成功", players, 1)
}
