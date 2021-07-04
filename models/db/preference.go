package db

import (
	cfg "go-netdisk/config"
	"time"
)

type Preference struct {
	UUID                  string    `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	Name                  string    `gorm:"column:name;type:varchar(255) not null" json:"name"`
	LogoURL               string    `gorm:"column:logo_url;type:varchar(255);default:''" json:"logoUrl"`
	FaviconURL            string    `gorm:"column:favicon_url;type:varchar(255);default:''" json:"faviconUrl"`
	Copyright             string    `gorm:"column:copyright;type:varchar(1024);comment:版权信息" json:"copyright"`
	Record                string    `gorm:"column:record;type:varchar(1024);comment:备案信息" json:"record"`
	DownloadDirMaxSize    int64     `gorm:"column:download_dir_max_size;not null;default:-1;comment:zip下载大小限制" json:"downloadDirMaxSize"`
	DownloadDirMaxNum     int64     `gorm:"column:download_dir_max_num;not null;default:-1;comment:zip下载数量限制" json:"downloadDirMaxNum"`
	DefaultTotalSizeLimit int64     `gorm:"column:default_total_size_limit;not null;default:-1;comment:用户默认总大小限制" json:"defaultTotalSizeLimit"`
	AllowRegister         bool      `gorm:"column:allow_register;not null;default:true;comment:允许自动注册" json:"allowRegister"`
	UpdateTime            time.Time `gorm:"column:update_time;not null" json:"updateTime"`
	CreateTime            time.Time `gorm:"column:create_time;not null" json:"createTime"`
}

func (Preference) TableName() string {
	return "preference"
}

func init() {
	// cfg.DB.AutoMigrate(&Preference{})
}

// Get preference record by uuid
func GetPreferenceByUUID(uuid string) (prefer *Preference, err error) {
	err = cfg.DB.First(&prefer, "uuid = ?", uuid).Error
	return
}
