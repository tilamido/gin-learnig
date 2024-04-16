package models

import (
	"social-network/dao"
	"time"
)

type User struct {
	Id         uint64    `gorm:"primaryKey" json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password,omitempty"`
	AddTime    time.Time `json:"add_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (User) TableName() string {
	return "users"
}

func GetUserInfoByName(username string) (User, error) {
	var user User
	err := dao.SqlDB.Where("username=?", username).First(&user).Error
	return user, err
}

func GetUserInfoByID(id uint64) (User, error) {
	var user User
	err := dao.SqlDB.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username, password string) (User, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.SqlDB.Create(&user).Error
	return user, err
}

func DeleteUser(id uint64) error {
	var user User
	err := dao.SqlDB.Where("id = ?", id).Delete(&user).Error
	return err
}

func ModifyPWD(id uint64, newpassword string) error {
	var user User
	err := dao.SqlDB.Model(&user).Where("id=?", id).Update("password", newpassword).Error
	return err
}

func GetAllUsers() ([]User, error) {
	var users []User
	if err := dao.SqlDB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetPageUserList(cnt, offset int) ([]User, error) {
	var users []User
	if err := dao.SqlDB.
		Order("add_time").
		Limit(cnt).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
