package db

import (
	"fmt"
	"go-netdisk/pkg/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// const mysqlDsn = "root:root@tcp(localhost:3306)/tank?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	dsn := Dsn()

	gormConfig := &gorm.Config{}
	if settings.ENV.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	mysqlDB, err := gorm.Open(mysql.Open(dsn), gormConfig)
	DB = mysqlDB

	return
}

func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.ENV.Mysql.Username,
		settings.ENV.Mysql.Password,
		settings.ENV.Mysql.Host,
		settings.ENV.Mysql.Port,
		settings.ENV.Mysql.Name,
	)
}
