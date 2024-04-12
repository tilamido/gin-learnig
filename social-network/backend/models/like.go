package models

import (
	"social-network/dao"
	"time"
)

type Like struct {
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	MomentID uint64    `gorm:"index;not null" json:"moment_id"`
	UserID   uint64    `gorm:"index;not null" json:"user_id"`
	AddTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"add_time"`
}

func AddLike(userid, momentid uint64) (*Like, error) {
	like := Like{
		MomentID: uint64(momentid),
		UserID:   uint64(userid),
	}
	if err := dao.SqlDB.Create(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}

func DelLike(userid, momentid uint64) (*Like, error) {
	like := Like{
		MomentID: uint64(momentid),
		UserID:   uint64(userid),
	}
	if err := dao.SqlDB.
		Where("user_id = ? AND moment_id = ?", like.UserID, like.MomentID).
		Delete(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}

func GetLikes(momentid uint64) ([]Like, error) {
	var like []Like
	if err := dao.SqlDB.Model(&like).Where("moment_id = ?", momentid).Error; err != nil {
		return nil, err
	}
	return like, nil

}
