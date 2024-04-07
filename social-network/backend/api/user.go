package api

import (
	"social-network/models"

	"github.com/gin-gonic/gin"
)

type UserAPI struct{}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

func (u UserAPI) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	username := req.Username
	password := req.Password

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
	ReturnSucess(c, 0, "登录成功", user.Username, 1)

}

func (u UserAPI) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	username := req.Username
	password := req.Password
	confirmPassword := req.ConfirmPassword

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

	ReturnSucess(c, 1, "注册成功", id, 1)

}
