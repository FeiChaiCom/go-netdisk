package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const MysqlDsn = "root:root@tcp(localhost:3306)/tank?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() error {
	if db, err := gorm.Open(mysql.Open(MysqlDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		return err
	} else {
		DB = db
	}
	return nil
}
