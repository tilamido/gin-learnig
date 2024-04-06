package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUser(id)
	ReturnSucess(c, 0, user.Username, user.Id, 1)
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "123456")
	id, err := models.AddUser(username, password)
	if err != nil {
		ReturnError(c, 402, "增加失败")
		return
	}
	ReturnSucess(c, 0, "增加成功", id, 1)
}

func (u UserController) GetList(c *gin.Context) {
	// logger.Write("日志", "user")

	num := 0
	ReturnError(c, 404, 1/num)

}
