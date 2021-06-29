package models

import (
	"database/sql"
	cfg "github.com/gaomugong/go-netdisk/config"
	"time"
)

type Matter struct { //nolint:maligned
	UUID       string         `gorm:"column:uuid;primaryKey;type:varchar(36)"`
	PUUID      string         `gorm:"column:puuid;type:varchar(36) not null"`
	Name       string         `gorm:"column:name;type:varchar(255) not null"`
	UserName   string         `gorm:"column:username;type:varchar(45) not null"`
	UserUUID   string         `gorm:"column:user_uuid;type:varchar(36) not null"`
	Md5        sql.NullString `gorm:"column:md5;type:varchar(45)"`
	Size       int            `gorm:"column:size;not null"`
	Dir        bool           `gorm:"column:dir;not null;default:false"`
	Privacy    bool           `gorm:"column:privacy;not null;default:true"`
	Path       string         `gorm:"column:path;type:varchar(1024)"`
	UpdateTime time.Time      `gorm:"column:update_time;not null"`
	CreateTime time.Time      `gorm:"column:create_time;not null"`
	Times      uint32         `gorm:"column:times;not null"`
	File       string         `gorm:"column:file;type:varchar(100) not null"`
}

func (Matter) TableName() string {
	return "matter"
}

func init() {
	cfg.DB.AutoMigrate(&Matter{})
}
