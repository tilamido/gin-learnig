package api

import (
	"social-network/models"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserAPI struct{}
type UserReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	NewPassword     string `json:"newpassword"`
}
type UserInfo struct {
	Id       uint64    `gorm:"primaryKey" json:"id"`
	Username string    `json:"username"`
	AddTime  time.Time `json:"add_time"`
}

func (u UserAPI) Login(c *gin.Context) {
	var req UserReq
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
	session.Set("login:"+strconv.FormatUint(user.Id, 10), user.Id)
	session.Save()

	userinfo := UserInfo{
		Id:       user.Id,
		Username: user.Username,
		AddTime:  user.AddTime,
	}
	ReturnSucess(c, 0, "登录成功", userinfo, 1)

}

func (u UserAPI) Register(c *gin.Context) {
	var req UserReq
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
	user, err := models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4005, "注册失败,联系管理")
		return
	}
	userinfo := UserInfo{
		Id:       user.Id,
		Username: user.Username,
		AddTime:  user.AddTime,
	}

	if user.Id != 0 {
		ReturnSucess(c, 0, "注册成功", userinfo, 1)
		return
	}

	ReturnError(c, 4005, "注册失败,联系管理")

}

func (u UserAPI) DeleteUser(c *gin.Context) {
	var req UserReq
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "参数错误")
		return
	}
	username := req.Username
	if username == "" {
		ReturnError(c, 4001, "参数错误")
		return
	}
	user, err := models.GetUserInfoByName(username)
	if err != nil {
		ReturnError(c, 4002, "查询失败")
		return
	}
	if user.Id == 0 {
		ReturnError(c, 4002, "用户不存在")
		return
	}
	if err := models.DeleteUser(user.Id); err != nil {
		ReturnError(c, 4003, "删除失败")
		return
	}

	userinfo := UserInfo{
		Id:       user.Id,
		Username: user.Username,
		AddTime:  user.AddTime,
	}
	ReturnSucess(c, 0, "删除成功", userinfo, 1)
}

func (u UserAPI) ModifyPWD(c *gin.Context) {
	var req UserReq
	if err := c.BindJSON(&req); err != nil {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	username := req.Username
	password := req.Password
	newpassword := req.NewPassword
	if username == "" || password == "" || newpassword == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

	if password == newpassword {
		ReturnError(c, 4001, "两次密码相同")
		return
	}
	user, err := models.GetUserInfoByName(username)
	if err != nil {
		ReturnError(c, 4001, "查询失败")
		return
	}
	if user.Id == 0 {
		ReturnError(c, 4002, "用户名不存在")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4003, "原始密码错误")
		return
	}

	err = models.ModifyPWD(user.Id, EncryMd5(newpassword))
	if err != nil {
		ReturnError(c, 4002, "修改失败")
		return
	}
	userinfo := UserInfo{
		Id:       user.Id,
		Username: user.Username,
		AddTime:  user.AddTime,
	}
	ReturnSucess(c, 0, "修改成功", userinfo, 1)

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
	ReturnSucess(c, 0, "查询到所有用户", usersinfo, int64(counts))

}

type PageReq struct {
	UserID uint64 `json:"user_id"`
	Counts int    `json:"counts"`
	Offset int    `json:"offset"`
}

func (u UserAPI) GetPageUserList(c *gin.Context) {
	var req PageReq
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

	ReturnSucess(c, 0, "查询每页用户", usersinfo, int64(len(usersinfo)))

}
