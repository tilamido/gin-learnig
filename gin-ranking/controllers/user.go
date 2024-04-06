package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserInfoByID(id)
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

func (u UserController) UpdateUserName(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	username := c.DefaultPostForm("username", "user")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(c, 403, "id错误")
	}
	if err := models.UpdateUserName(id, username); err != nil {
		ReturnError(c, 403, "修改失败")
		return
	}
	ReturnSucess(c, 0, "修改成功", id, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ReturnError(c, 403, "id错误")
	}
	if err := models.DeleteUser(id); err != nil {
		ReturnError(c, 403, "删除失败")
		return
	}
	ReturnSucess(c, 0, "删除成功", id, 1)
}

// func (u UserController) GetList(c *gin.Context) {
// 	// logger.Write("日志", "user")

// 	num := 0
// 	ReturnError(c, 404, 1/num)

// }

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")

	password := c.DefaultPostForm("password", "")

	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "两次密码不同")
		return
	}

	user, _ := models.GetUserInfoByName(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}

	id, err := models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4002, "注册失败,联系管理")
		return
	}

	ReturnSucess(c, 1, id, "注册成功", 1)

}

type UserApi struct {
	ID       int    `josn:"id"`
	Username string `josn:"username"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	user, _ := models.GetUserInfoByName(username)
	if user.Id == 0 {
		ReturnError(c, 4001, "用户名不存在")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "密码错误")
		return
	}
	UserApi := UserApi{
		Username: user.Username,
		ID:       user.Id,
	}
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(UserApi.ID), UserApi.ID)
	session.Save()

	data := UserApi
	ReturnSucess(c, 0, "登录成功", data, 1)

}
