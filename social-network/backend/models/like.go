package models

import (
	"social-network/dao"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := dao.SqlDB.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}

func DelLike(userid, momentid uint64) error {
	if err := dao.SqlDB.
		Where("user_id = ? AND moment_id = ?", userid, momentid).
		Delete(&Like{}).Error; err != nil {
		return err
	}
	return nil
}

func GetLikes(momentid uint64) ([]Like, error) {
	var likes []Like
	if err := dao.SqlDB.Where("moment_id = ?", momentid).Find(&likes).Error; err != nil {
		return nil, err
	}
	return likes, nil
}

func CheckLikeStatus(momentID, userID uint64) (bool, error) {
	var like Like
	err := dao.SqlDB.Where("moment_id = ? AND user_id = ?", momentID, userID).First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
