package models

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
	"math"
	"time"
)

type Matter struct { //nolint:maligned
	UUID       string    `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	PUUID      string    `gorm:"column:puuid;type:varchar(36) not null" json:"puuid"`
	Name       string    `gorm:"column:name;type:varchar(255) not null" json:"name"`
	UserName   string    `gorm:"column:username;type:varchar(45) not null" json:"username"`
	UserUUID   string    `gorm:"column:user_uuid;type:varchar(36) not null" json:"userUuid"`
	Md5        string    `gorm:"column:md5;type:varchar(45)" json:"md5"`
	Size       int       `gorm:"column:size;not null" json:"size"`
	Dir        bool      `gorm:"column:dir;not null;default:false" json:"dir"`
	Privacy    bool      `gorm:"column:privacy;not null;default:true" json:"privacy"`
	Path       string    `gorm:"column:path;type:varchar(1024)" json:"path"`
	UpdateTime time.Time `gorm:"column:update_time;not null" json:"updateTime"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"createTime"`
	Times      uint32    `gorm:"column:times;not null" json:"times"`
	File       string    `gorm:"column:file;type:varchar(100) not null" json:"file"`
}

func (Matter) TableName() string {
	return "matter"
}

func init() {
	// cfg.DB.AutoMigrate(&Matter{})
}

func GetAllMatters(puuid string, name string, page int, pageSize int, order string) (matters []*Matter, totalItems int64, totalPage int) {
	// cfg.DB.Model(&Matter{}).Where(Matter{PUUID: puuid}).Count(&totalItems)
	tx := cfg.DB.Model(&Matter{}).Where("puuid = ?", puuid)

	if name != "" {
		tx = tx.Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", name, "%"))
	}

	tx.Count(&totalItems)

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	totalPage = int(math.Round(float64((totalItems + int64(pageSize) - 1) / int64(pageSize))))
	if page > totalPage {
		page = 1
	}

	if page > 0 {
		if offset := (page - 1) * pageSize; offset > 0 {
			tx = tx.Offset(offset)
		}
	}

	if order != "" {
		order = fmt.Sprintf("create_time %s", order)
		tx = tx.Order(order)
	}

	tx.Limit(pageSize).Find(&matters)

	return
}
