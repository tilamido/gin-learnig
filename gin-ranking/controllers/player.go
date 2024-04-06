package controllers

import (
	"gin-ranking/cache"
	"gin-ranking/models"
	"strconv"
	"time"

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

func (p PlayerController) GetRankingByRedis(c *gin.Context) {

	aidstr := c.DefaultPostForm("aid", "-1")
	aid, err := strconv.Atoi(aidstr)
	if aid == -1 || err != nil {
		ReturnError(c, 0, "活动不存在")
		return
	}

	redisKey := "ranking:" + aidstr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayersByID(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSucess(c, 0, "success from cache", players, 1)
		return
	}
	rssqlDb, errsqlDb := models.GetPlayersRankByAID(aid, "score desc")
	if errsqlDb == nil {
		for _, value := range rssqlDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score))
		}
		cache.Rdb.Expire(cache.Rctx, redisKey, time.Hour*24)
		ReturnSucess(c, 0, "success from sql", rssqlDb, 1)
		return
	}
	ReturnError(c, 401, "查询失败")
}
