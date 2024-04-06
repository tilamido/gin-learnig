package models

import (
	"gin-ranking/dao"
	"time"
)

type Vote struct {
	Id       int       `gorm:"primaryKey" json:"id"`
	UserID   int       `gorm:"index;not null" json:"user_id"`
	PlayerID int       `gorm:"index;not null" json:"player_id"`
	AddTime  time.Time `json:"add_time"`
}

func (Vote) TableName() string {
	return "votes"
}

func GetVoteInfo(uID int, pID int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id = ? AND player_id = ?", uID, pID).Error
	return vote, err
}

func AddVote(uID int, pID int) (int, error) {
	vote := Vote{
		UserID:   uID,
		PlayerID: pID,
		AddTime:  time.Now(),
	}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
