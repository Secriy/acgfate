package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	Male = iota
	Female
	Private
	Other
)

var GenderFlags = map[int]string{
	Male:    "男",
	Female:  "女",
	Private: "保密",
	Other:   "其他",
}

type User struct {
	UID       uint64 `gorm:"primaryKey;unique;autoIncrement;comment:'用户ID'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;comment:'用户名'"`
	Password  string         `gorm:"comment:'密码'" `
	Nickname  string         `gorm:"comment:'昵称'"`
	Mail      string         `gorm:"comment:'邮箱'"`
	Avatar    string         `gorm:"comment:'头像'"`
	Gender    uint8          `gorm:"comment:'性别'"`
	Birthday  time.Time      `gorm:"comment:'生日'"`
	Level     uint8          `gorm:"comment:'等级'"`
	JoinTime  time.Time      `gorm:"comment:'加入时间'"`
	Silence   bool           `gorm:"comment:'禁言'"`
}

// CheckPass 检查密码是否正确
func (u *User) CheckPass(password string) bool {
	if u.Password == password {
		return true
	}

	return false
}

// GetUser 获取当前用户模型
func GetUser(uid interface{}) (User, error) {
	var user User
	res := DB.First(&user, uid)

	return user, res.Error
}
