package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MYSQL_DSN = "root:88888888@tcp(localhost:3306)/tank?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() error {
	if db, err := gorm.Open(mysql.Open(MYSQL_DSN), &gorm.Config{}); err != nil {
		return err
	} else {
		DB = db
	}
	return nil
}
