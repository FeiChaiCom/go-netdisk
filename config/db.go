package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const MysqlDsn = "root:88888888@tcp(localhost:3306)/tank?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() (err error) {
	gormConfig := &gorm.Config{}
	if DebugOn {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(mysql.Open(MysqlDsn), gormConfig)
	if err != nil {
		return
	}
	return
}
