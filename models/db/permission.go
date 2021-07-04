package db

import (
	cfg "go-netdisk/config"
	"time"
)

// Role types
const (
	USER          = "USER"
	PROJECT_ADMIN = "PROJECT_ADMIN"
	ADMINISTRATOR = "ADMINISTRATOR"
)

type Permission struct {
	UUID        string    `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	UserName    string    `gorm:"column:username;type:varchar(45) not null" json:"username"`
	ProjectUUID string    `gorm:"column:project_uuid;type:varchar(36) not null" json:"projectUuid"`
	Role        string    `gorm:"column:role;type:varchar(45);default:USER" json:"role"`
	CreateAt    time.Time `gorm:"column:create_at;not null" json:"createAt"`
	CreateBy    string    `gorm:"column:create_by;type:varchar(45) not null" json:"createBy"`
	UpdateAt    time.Time `gorm:"column:update_at" json:"updateAt"`
	UpdateBy    string    `gorm:"column:create_by;type:varchar(45)" json:"updateBy"`
	IsDeleted   bool      `gorm:"column:is_deleted;not null;default:false" json:"isDeleted"`
	DeletedAt   time.Time `gorm:"column:update_at" json:"DeletedAt"`
	DeletedBy   string    `gorm:"column:create_by;type:varchar(45)" json:"DeletedBy"`
}

func (Permission) TableName() string {
	return "permission"
}

func init() {
	// cfg.DB.AutoMigrate(&Permission{})
}

// Get user's permission
func GetPermissionByUsername(username string) (permission *Permission, err error) {
	err = cfg.DB.First(&permission, "username = ?", username).Error
	return
}
