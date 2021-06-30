package models

import (
	"fmt"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"math"
	"time"
)

type User struct {
	UUID           uuid.UUID `gorm:"column:uuid;primaryKey;type:varchar(32)" json:"uuid"`
	Username       string    `gorm:"column:username;type:varchar(45) not null;unique" json:"username"`
	Password       string    `gorm:"column:password;type:varchar(255) not null" json:"-"`
	Role           string    `gorm:"column:role;type:varchar(45);default:USER" json:"role"`
	Status         string    `gorm:"column:status;not null;default:OK" json:"status"`
	SizeLimit      int64     `gorm:"column:size_limit;not null;default:268435456" json:"sizeLimit"`
	TotalSizeLimit int64     `gorm:"column:total_size_limit;not null;default:-1" json:"totalSizeLimit"`
	TotalSize      int64     `gorm:"column:total_size;not null;default:0" json:"totalSize"`
	AvatarURL      string    `gorm:"column:avatar_url;type:varchar(255)" json:"avatarUrl"`
	UpdateTime     time.Time `gorm:"column:update_time;not null" json:"updateTime"`
	CreateTime     time.Time `gorm:"column:create_time;not null" json:"createTime"`
	LastTime       time.Time `gorm:"column:last_time;not null" json:"lastTime"`
	LastIP         string    `gorm:"column:last_ip;type:varchar(128)" json:"lastIp"`
}

func (User) TableName() string {
	return "accounts_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.NewV4()
	return nil
}

func init() {
	// cfg.DB.AutoMigrate(&User{})
}

func GetAllUsers(page int, pageSize int, order string) (users []*User, totalItems int64, totalPage int) {
	tx := cfg.DB.Model(&User{})
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

	tx.Limit(pageSize).Find(&users)

	return
}
