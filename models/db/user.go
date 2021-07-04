package db

import (
	"errors"
	"fmt"
	"github.com/alexandrevicenzi/unchained"
	cfg "github.com/gaomugong/go-netdisk/config"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"math"
	"time"
)

type User struct {
	UUID           uuid.UUID `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	Username       string    `gorm:"column:username;type:varchar(45) not null;unique" json:"username"`
	Password       string    `gorm:"column:password;type:varchar(255) not null" json:"-"`
	IsSuperUser    bool      `gorm:"column:is_superuser;default:false" json:"-"`
	IsStaff        bool      `gorm:"column:is_staff;default:false" json:"-"`
	IsActive       bool      `gorm:"column:is_active;default:false" json:"-"`
	FirstName      string    `gorm:"column:first_name;type:varchar(30);default:''"`
	LastName       string    `gorm:"column:last_name;type:varchar(150);default:''"`
	Email          string    `gorm:"column:email;default:''"`
	Role           string    `gorm:"column:role;type:varchar(45);default:USER" json:"role"`
	Status         string    `gorm:"column:status;not null;default:OK" json:"status"`
	SizeLimit      int64     `gorm:"column:size_limit;not null;default:268435456" json:"sizeLimit"`
	TotalSizeLimit int64     `gorm:"column:total_size_limit;not null;default:-1" json:"totalSizeLimit"`
	TotalSize      int64     `gorm:"column:total_size;not null;default:0" json:"totalSize"`
	AvatarURL      string    `gorm:"column:avatar_url;type:varchar(255)" json:"avatarUrl"`
	UpdateTime     time.Time `gorm:"column:update_time;not null" json:"updateTime"`
	CreateTime     time.Time `gorm:"column:create_time;not null" json:"createTime"`
	LastTime       time.Time `gorm:"column:last_time;default:null" json:"lastTime"`
	DateJoined     time.Time `gorm:"column:date_joined;default:null" json:"dateJoined"`
	LastIP         string    `gorm:"column:last_ip;type:varchar(128)" json:"lastIp"`
}

type LoginParam struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterParam struct {
	LoginParam
}

func (User) TableName() string {
	return "accounts_user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.NewV4()
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	u.DateJoined = time.Now()
	u.LastTime = time.Now()
	return nil
}

func init() {
	// cfg.DB.AutoMigrate(&User{})
}

var ErrUserExist = errors.New("username is registered, please replace another one")

func GetUserByUUID(uuid string) (user *User, err error) {
	err = cfg.DB.First(&user, "uuid = ?", uuid).Error
	return
}

func GetUserByName(username string) (user *User, err error) {
	err = cfg.DB.Where("username = ?", username).First(&user).Error
	return
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

func Register(u User) (user *User, err error) {
	err = cfg.DB.Where("username = ?", u.Username).First(&user).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, ErrUserExist
	}

	u.Password, err = unchained.MakePassword(u.Password, "", unchained.Argon2Hasher)
	if err != nil {
		errMsg := fmt.Sprintf("register failed: create password error, %s", err.Error())
		return user, errors.New(errMsg)
	}

	err = cfg.DB.Create(&u).Error

	return user, err
}

func Login(p *LoginParam) (u *User, err error) {
	if u, err = GetUserByName(p.Username); err != nil {
		return u, errors.New("invalid user")
	}

	if isValid, err := unchained.CheckPassword(p.Password, u.Password); err != nil || !isValid {
		return u, errors.New("invalid user or password")
	}

	return
}
