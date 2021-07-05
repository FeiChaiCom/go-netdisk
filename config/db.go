package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// const mysqlDsn = "root:root@tcp(localhost:3306)/tank?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ENV.Mysql.Username,
		ENV.Mysql.Password,
		ENV.Mysql.Host,
		ENV.Mysql.Port,
		ENV.Mysql.Name,
	)

	gormConfig := &gorm.Config{}
	if ENV.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	return
}
