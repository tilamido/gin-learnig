package models

import "time"

type Activity struct {
	Id      int       `gorm:"primaryKey" json:"id"`
	Name    string    `json:"name"`
	AddTime time.Time `json:"add_time"`
}

func (Activity) TableName() string {
	return "activities"
}
