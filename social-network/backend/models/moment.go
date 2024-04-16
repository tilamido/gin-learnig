package models

import (
	"social-network/dao"
	"time"
)

type Moment struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID     uint64    `gorm:"index;notNull;column:user_id" json:"user_id"`
	Content    string    `gorm:"type:text;notNull;column:content" json:"content"`
	ImagePaths string    `gorm:"type:text;column:image_paths" json:"image_paths"`
	AddTime    time.Time `gorm:"autoCreateTime;column:add_time" json:"add_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime;column:update_time" json:"update_time"`
}

func (Moment) TableName() string {
	return "moments"
}

func AddMoment(userid uint64, content, imgpath string) error {
	moment := Moment{
		UserID:     userid,
		Content:    content,
		ImagePaths: imgpath,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	err := dao.SqlDB.Create(&moment).Error
	return err
}

func GetPageMomentByTime(cnt int, offset int) ([]Moment, error) {
	var moments []Moment
	if err := dao.SqlDB.Order("update_time DESC").Limit(cnt).Offset(offset).Find(&moments).Error; err != nil {
		return nil, err
	}
	return moments, nil
}

func GetAllMomentByTime(cnt int, offset int) ([]Moment, error) {
	var moments []Moment
	if err := dao.SqlDB.Order("update_time DESC").Limit(cnt).Offset(offset).Find(&moments).Error; err != nil {
		return nil, err
	}
	return moments, nil
}

func GetUserMoments(userid uint64, cnt int, offset int) ([]Moment, error) {
	var moments []Moment
	if err := dao.SqlDB.Debug().Where("user_id = ?", userid).Order("update_time desc").Limit(cnt).Offset(offset).Find(&moments).Error; err != nil {
		return nil, err
	}
	return moments, nil
}

func DeleteMoment(momentID uint64) error {
	if err := dao.SqlDB.Where("id = ?", momentID).Delete(&Moment{}).Error; err != nil {
		return err
	}
	return nil
}

func GetMomentByID(id uint64) (Moment, error) {
	var moment Moment
	err := dao.SqlDB.First(&moment, id).Error
	return moment, err
}

type LikeInfo struct {
	MomentID   uint64
	Likescount uint64
}

func RankMoments(cnt, offset int) ([]LikeInfo, error) {
	var likeinfos []LikeInfo
	if err := dao.SqlDB.Table("likes").
		Select("moment_id,COUNT(*) AS likes_count").
		Group("moment_id").
		Order("likes_count DESC").
		Limit(cnt).
		Offset(offset).
		Scan(&likeinfos).
		Error; err != nil {
		return nil, err
	}
	return likeinfos, nil
}
