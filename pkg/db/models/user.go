package models

import (
	"errors"
	"fmt"
	"github.com/alexandrevicenzi/unchained"
	"github.com/satori/go.uuid"
	"go-netdisk/pkg/db"
	"go-netdisk/pkg/settings"
	"gorm.io/gorm"
	"math"
	"time"
)

type User struct {
	UUID           uuid.UUID `gorm:"column:uuid;primaryKey;type:varchar(36)" json:"uuid"`
	Username       string    `gorm:"column:username;type:varchar(45) not null;unique" json:"username"`
	Password       string    `gorm:"column:password;type:varchar(255) not null" json:"-"`
	Role           string    `gorm:"column:role;type:varchar(45);default:USER" json:"role"`
	Status         string    `gorm:"column:status;not null;default:OK" json:"status"`
	SizeLimit      int64     `gorm:"column:size_limit;not null;default:268435456" json:"sizeLimit"`
	TotalSizeLimit int64     `gorm:"column:total_size_limit;not null;default:-1" json:"totalSizeLimit"`
	TotalSize      int64     `gorm:"column:total_size;not null;default:0" json:"totalSize"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`
	LastTime       time.Time `gorm:"column:last_time" json:"lastTime"`
	LastIP         string    `gorm:"column:last_ip;type:varchar(128)" json:"lastIp"`
	// Django specific user column
	DateJoined  time.Time `gorm:"column:date_joined" json:"-"`
	AvatarURL   string    `gorm:"column:avatar_url;type:varchar(255)" json:"-"`
	IsSuperUser bool      `gorm:"column:is_superuser;default:false" json:"-"`
	IsStaff     bool      `gorm:"column:is_staff;default:false" json:"-"`
	IsActive    bool      `gorm:"column:is_active;default:false" json:"isActive"`
	FirstName   string    `gorm:"column:first_name;type:varchar(30);default:''" json:"-"`
	LastName    string    `gorm:"column:last_name;type:varchar(150);default:''" json:"-"`
	Email       string    `gorm:"column:email;default:''" json:"-"`
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

var ErrUserExist = errors.New("username is registered, please replace another one")

// GetOrCreateUser get or create user by username
func GetOrCreateUser(username string, isSuperUser bool) (user *User, err error) {
	user, err = GetUserByName(username)

	// User exist, return directly
	if err == nil {
		return user, nil
	}

	// Return error except not found
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, err
	}

	// Not found, then create a new user
	password, err := unchained.MakePassword(settings.ENV.DefaultPassword, "", unchained.Argon2Hasher)
	if err != nil {
		return user, err
	}

	role := "USER"
	if isSuperUser {
		role = "ADMINISTRATOR"
	}

	user = &User{
		Username: username,
		Password: password,
		Role:     role,
	}

	err = db.DB.Create(user).Error

	return
}

func GetUserByUUID(uuid string) (user *User, err error) {
	err = db.DB.First(&user, "uuid = ?", uuid).Error
	return
}

func GetUserByName(username string) (user *User, err error) {
	err = db.DB.Where("username = ?", username).First(&user).Error
	return
}

func GetAllUsers(page int, pageSize int, order string) (users []*User, totalItems int64, totalPage int) {
	tx := db.DB.Model(&User{})
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

// Register a new user with username and password
func Register(r *RegisterParam) (user *User, err error) {
	user, err = GetUserByName(r.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return user, ErrUserExist
	}

	password, err := unchained.MakePassword(r.Password, "", unchained.Argon2Hasher)
	if err != nil {
		errMsg := fmt.Sprintf("register failed: create password error, %s", err.Error())
		return user, errors.New(errMsg)
	}

	user = &User{
		Username: r.Username,
		Password: password,
	}
	err = db.DB.Create(user).Error

	return
}

// Login check user exist and password match username
func Login(p *LoginParam) (u *User, err error) {
	if u, err = GetUserByName(p.Username); err != nil {
		return u, errors.New("invalid user")
	}

	isValid, err := unchained.CheckPassword(p.Password, u.Password)
	if err != nil || !isValid {
		return u, errors.New("invalid user or password")
	}

	// log.Printf("check password: %s vs %#v -> %#v, %#v", p.Password, u.Password, isValid, err)
	return
}
