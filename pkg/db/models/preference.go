package models

import (
	uuid "github.com/satori/go.uuid"
	"go-netdisk/pkg/db"
	"gorm.io/gorm"
	"time"
)

type Preference struct {
	UUID                  uuid.UUID `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	Name                  string    `gorm:"column:name;type:varchar(255) not null" json:"name"`
	LogoURL               string    `gorm:"column:logo_url;type:varchar(255);default:''" json:"logoUrl"`
	FaviconURL            string    `gorm:"column:favicon_url;type:varchar(255);default:''" json:"faviconUrl"`
	Copyright             string    `gorm:"column:copyright;type:varchar(1024);comment:版权信息" json:"copyright"`
	Record                string    `gorm:"column:record;type:varchar(1024);comment:备案信息" json:"record"`
	DownloadDirMaxSize    int64     `gorm:"column:download_dir_max_size;not null;default:-1;comment:zip下载大小限制" json:"downloadDirMaxSize"`
	DownloadDirMaxNum     int64     `gorm:"column:download_dir_max_num;not null;default:-1;comment:zip下载数量限制" json:"downloadDirMaxNum"`
	DefaultTotalSizeLimit int64     `gorm:"column:default_total_size_limit;not null;default:-1;comment:用户默认总大小限制" json:"defaultTotalSizeLimit"`
	AllowRegister         bool      `gorm:"column:allow_register;not null;default:true;comment:允许自动注册" json:"allowRegister"`
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`
}

func (Preference) TableName() string {
	return "preference"
}

func (obj *Preference) BeforeCreate(tx *gorm.DB) (err error) {
	obj.UUID = uuid.NewV4()
	obj.CreateTime = time.Now()
	obj.UpdateTime = time.Now()
	return nil
}

// Get preference record by uuid
func GetPreferenceByUUID(uuid string) (prefer *Preference, err error) {
	err = db.DB.First(&prefer, "uuid = ?", uuid).Error
	return
}
