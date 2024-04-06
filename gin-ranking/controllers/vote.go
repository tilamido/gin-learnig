package controllers

import (
	"gin-ranking/cache"
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	userIDstr := c.DefaultPostForm("userID", "0")
	playerIDstr := c.DefaultPostForm("playerID", "0")
	userID, _ := strconv.Atoi(userIDstr)
	playerID, _ := strconv.Atoi(playerIDstr)
	if userID == 0 || playerID == 0 {
		ReturnError(c, 4001, "请输入准确信息")
		return
	}
	user, _ := models.GetUserInfoByID(userID)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}
	player, _ := models.GetPlayersByID(playerID)
	if player.Id == 0 {
		ReturnError(c, 4001, "参赛选手不存在")
		return
	}
	Vote, _ := models.GetVoteInfo(user.Id, player.Id)
	if Vote.Id != 0 {
		ReturnError(c, 4001, "已投过票")
		return
	}
	rs, err := models.AddVote(user.Id, player.Id)
	if err == nil {
		models.UpdatePlayerScore(player.Id)
		redisKey := "ranking:" + strconv.Itoa(player.ActivityId)
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, playerIDstr)
		ReturnSucess(c, 0, "投票成功", rs, 1)
		return
	}
	ReturnError(c, 0, "操作失败")

}
