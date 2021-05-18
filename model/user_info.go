package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	EncryptCost = 12 // 加密强度
)

const (
	Private = iota + 1 // 性别: 保密
	Male               // 性别: 男
	Female             // 性别: 女
	Other              // 性别: 其他
)

var GenderFlags = map[int]string{
	Private: "保密",
	Male:    "男",
	Female:  "女",
	Other:   "其他",
}

type User struct {
	UserInfo
	UserPoints
}

type UserInfo struct {
	UID       uint64 `gorm:"primaryKey;unique;autoIncrement;comment:'用户ID'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;comment:'用户名'"`
	Password  string         `gorm:"comment:'密码'" `
	Nickname  string         `gorm:"comment:'昵称'"`
	Mail      string         `gorm:"comment:'邮箱'"`
	Avatar    string         `gorm:"comment:'头像'"`
	Gender    uint8          `gorm:"comment:'性别';default:1"`
	Birthday  string         `gorm:"comment:'生日'"`
	JoinTime  time.Time      `gorm:"comment:'加入时间'"`
	Silence   bool           `gorm:"comment:'禁言';default:false"`
}

// type PremiumUser struct {
// }

// SetPassword 设置密码
func (info *UserInfo) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), EncryptCost)
	if err != nil {
		return err
	}
	info.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (info *UserInfo) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password))
	return err == nil
}

// GetUser 获取当前用户模型
func GetUser(uid interface{}) (user User, err error) {
	var userInfo UserInfo
	var userPoints UserPoints
	err = DB.First(&userInfo, uid).Error
	if err != nil {
		return
	}
	err = DB.First(&userPoints, uid).Error
	user = User{
		userInfo,
		userPoints,
	}

	return
}
