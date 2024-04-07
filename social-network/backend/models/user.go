package models

import (
	"social-network/dao"
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

func GetUserInfoByName(username string) (User, error) {
	var user User
	err := dao.SqlDB.Where("username=?", username).First(&user).Error
	return user, err
}
func AddUser(username, password string) (int, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.SqlDB.Create(&user).Error
	return user.Id, err
}
