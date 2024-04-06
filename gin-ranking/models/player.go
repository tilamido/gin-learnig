package models

import (
	"gin-ranking/dao"
	"time"

	"gorm.io/gorm"
)

type Player struct {
	Id               int       `gorm:"primaryKey" json:"id"`
	ActivityId       int       `gorm:"index;not null" json:"activity_id"`
	ContestantNumber string    `json:"contestant_number"`
	Nickname         string    `json:"nickname"`
	Description      string    `gorm:"type:text" json:"description,omitempty"`
	Image            string    `json:"image,omitempty"`
	Score            float64   `gorm:"type:decimal(10,2)" json:"score"`
	AddTime          time.Time `json:"add_time"`
	UpdateTime       time.Time `json:"update_time"`
}

func (Player) TableName() string {
	return "players"
}

func GetPlayersByAID(aid int) ([]Player, error) {
	var players []Player
	err := dao.Db.Model(&players).Where("activity_id = ?", aid).Find(&players).Error
	return players, err
}

func GetPlayersByID(aid int) (Player, error) {
	var player Player
	err := dao.Db.Model(&player).Where("id = ?", aid).First(&player).Error
	return player, err
}

func UpdatePlayerScore(pID int) error {
	var player Player
	err := dao.Db.Model(&player).Where("id = ?", pID).UpdateColumn("score", gorm.Expr("score + ?", 1)).Error
	return err

}

func GetPlayersRankByAID(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Model(&players).Where("activity_id = ?", aid).Order(sort).Find(&players).Error
	return players, err
}
