package dao

import (
	"social-network/config"
	"social-network/middleware/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	SqlDB *gorm.DB
	err   error
)

func init() {
	SqlDB, err = gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{})

	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}

	db, err := SqlDB.DB()
	if err != nil {
		logger.Error(map[string]interface{}{"datebase error": err.Error()})
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
}
