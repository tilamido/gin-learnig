package models

import (
	"fmt"
	"gin-ranking/dao"
	"time"
)

type User struct {
	Id         int       `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password,omitempty"`
	AddTime    time.Time `json:"add_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (User) TableName() string {
	return "users"
}

func GetUserInfoByID(id int) (User, error) {
	var user User
	err := dao.Db.Where("id =?", id).First(&user).Error
	return user, err
}
func GetUserInfoByName(username string) (User, error) {
	var user User
	err := dao.Db.Where("username =?", username).First(&user).Error
	return user, err
}
func AddUser(username, password string) (int, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
func GetAllUsers() ([]User, error) {
	var users []User
	err := dao.Db.Find(&users).Error
	return users, err
}
func UpdateUserName(id int, newUsername string) error {
	fmt.Print(newUsername)
	var user User
	err := dao.Db.Model(&user).Where("id = ?", id).Update("username", newUsername).Error
	return err
}
func UpdateUserPwd(id int, newUserPwd string) error {
	var user User
	err := dao.Db.Model(&user).Where("id = ?", id).Update("password", newUserPwd).Error
	return err
}
func DeleteUser(id int) error {
	var user User
	err := dao.Db.Delete(&user, id).Error
	return err
}
