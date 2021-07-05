package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// const mysqlDsn = "root:root@tcp(localhost:3306)/tank?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() (err error) {
	m := viper.GetStringMapString("mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m["username"],
		m["password"],
		m["host"],
		m["port"],
		m["name"],
	)

	gormConfig := &gorm.Config{}
	if viper.GetBool("debug") {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	return
}
