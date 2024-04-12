package api

import (
	"social-network/models"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserAPI struct{}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()

	ReturnSucess(c, 0, "登录成功", user.Username, 1)

}

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
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
		ReturnError(c, 4002, "请输入正确信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4003, "两次密码不同")
		return
	}
	user, _ := models.GetUserInfoByName(username)
	if user.Id != 0 {
		ReturnError(c, 4004, "用户名已存在")
		return
	}
	id, err := models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4005, "注册失败,联系管理")
		return
	}
	if id != 0 {
		ReturnSucess(c, 2000, "注册成功", id, 1)

		return
	}
	ReturnError(c, 4005, "注册失败,联系管理")

}

func (u UserAPI) DeleteUser(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}

	if _, err := models.GetUserInfoByName(req.Username); err != nil {
		ReturnError(c, 4002, "用户不存在")
		return
	}
	if err := models.DeleteUser(req.Username); err != nil {
		ReturnError(c, 4002, "删除失败")
		return
	}

	ReturnSucess(c, 2000, "删除成功", req, 1)
}

type ModifyRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

func (u UserAPI) ModifyPWD(c *gin.Context) {
	var req ModifyRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	username := req.Username
	password := req.Password
	newpassword := req.NewPassword

	if password == newpassword {
		ReturnError(c, 4001, "两次密码相同")
		return
	}
	user, _ := models.GetUserInfoByName(username)
	if user.Id == 0 {
		ReturnError(c, 4002, "用户名不存在")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4003, "原始密码错误")
		return
	}

	err := models.ModifyPWD(username, EncryMd5(newpassword))
	if err != nil {
		ReturnError(c, 4002, "修改失败")
		return
	}

	ReturnSucess(c, 2000, "修改成功", username, 1)

}

type UserInfo struct {
	Id       int       `gorm:"primaryKey" json:"id"`
	Username string    `json:"username"`
	AddTime  time.Time `json:"add_time"`
}

func (u UserAPI) GetAllUserList(c *gin.Context) {

	users, err := models.GetAllUsers()
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	var usersinfo []UserInfo
	for _, user := range users {
		userinfo := UserInfo{
			Id:       user.Id,
			Username: user.Username,
			AddTime:  user.AddTime,
		}
		usersinfo = append(usersinfo, userinfo)

	}
	counts := len(usersinfo)
	ReturnSucess(c, 2000, "查询到所有用户", usersinfo, int64(counts))

}

type PageListRequest struct {
	Counts int `json:"counts"`
	Offset int `json:"offset"`
}

func (u UserAPI) GetPageUserList(c *gin.Context) {
	var req PageListRequest
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数有误")
		return
	}
	counts := req.Counts
	offset := req.Offset
	users, err := models.GetPageUserList(counts, offset)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	var usersinfo []UserInfo
	for _, user := range users {
		userinfo := UserInfo{
			Id:       user.Id,
			Username: user.Username,
			AddTime:  user.AddTime,
		}
		usersinfo = append(usersinfo, userinfo)

	}

	ReturnSucess(c, 2000, "查询每页用户", usersinfo, int64(counts))

}
