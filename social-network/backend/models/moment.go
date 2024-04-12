package models

import (
	"social-network/dao"
	"time"
)

type Moment struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID     uint64    `gorm:"index;notNull;column:user_id" json:"user_id"`
	Username   string    `gorm:"-" json:"username"`
	Content    string    `gorm:"type:text;notNull;column:content" json:"content"`
	ImagePaths string    `gorm:"type:text;column:image_paths" json:"image_paths"`
	AddTime    time.Time `gorm:"autoCreateTime;column:add_time" json:"add_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"update_time"`
}

func (Moment) TableName() string {
	return "moments"
}

func AddMoment(userid uint64, content, imgpath string) (uint64, error) {
	moment := Moment{
		UserID:     userid,
		Content:    content,
		ImagePaths: imgpath,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.SqlDB.Create(&moment).Error
	return moment.ID, err

}

func GetPageMomentListByTime(cnt int, offset int) ([]Moment, error) {
	var moments []Moment
	if err := dao.SqlDB.Table("moments").Select("moments.*,users.username").
		Joins("left join users on users.id = moments.user_id").
		Order("moments.update_time desc").
		Limit(cnt).
		Offset(offset).
		Scan(&moments).Error; err != nil {
		return nil, err
	}
	return moments, nil
}

func GetAllMomentListByTime() ([]Moment, error) {
	var moments []Moment
	if err := dao.SqlDB.Table("moments").Select("moments.*,users.username").
		Joins("left join users on users.id = moments.user_id").
		Order("moments.update_time desc").
		Scan(&moments).Error; err != nil {
		return nil, err
	}
	return moments, nil
}

func GetUserMoments(userid uint64) ([]Moment, error) {
	var moments []Moment
	if err := dao.SqlDB.Order("update_time desc").Find(&moments).Error; err != nil {
		return nil, err
	}
	return moments, nil
}

func DeleteMoment(momentID uint64) error {
	var moment Moment
	if err := dao.SqlDB.Where("id = ?", momentID).Delete(&moment).Error; err != nil {
		return err
	}
	return nil
}
