package db

import (
	cfg "github.com/gaomugong/go-netdisk/config"
	"time"
)

type Project struct {
	UUID        string    `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	Name        string    `gorm:"column:name;type:varchar(64) not null;comment:项目名称" json:"name"`
	Description string    `gorm:"column:description;type:varchar(255);default:'';comment:项目描述" json:"description"`
	CreateAt    time.Time `gorm:"column:create_at;not null" json:"createAt"`
	CreateBy    string    `gorm:"column:create_by;type:varchar(45) not null" json:"createBy"`
	UpdateAt    time.Time `gorm:"column:update_at" json:"updateAt"`
	UpdateBy    string    `gorm:"column:create_by;type:varchar(45)" json:"updateBy"`
	IsDeleted   bool      `gorm:"column:is_deleted;not null;default:false" json:"isDeleted"`
	DeletedAt   time.Time `gorm:"column:update_at" json:"DeletedAt"`
	DeletedBy   string    `gorm:"column:create_by;type:varchar(45)" json:"DeletedBy"`
}

func (Project) TableName() string {
	return "project"
}

func init() {
	// cfg.DB.AutoMigrate(&Project{})
}

// Get user's project
func GetProjectByUsername(username string) (project *Project, err error) {
	err = cfg.DB.First(&project, "username = ?", username).Error
	return
}
