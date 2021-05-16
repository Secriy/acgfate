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

type User struct {
	UID       uint64 `gorm:"primaryKey;unique;autoIncrement;comment:'用户ID'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"unique;comment:'用户名'" json:"username"`
	Password  string         `gorm:"comment:'密码'" json:"password"`
	Nickname  string         `gorm:"comment:'昵称'" json:"nickname"`
	Mail      string         `gorm:"comment:'邮箱'" json:"mail"`
	Avatar    string         `gorm:"comment:'头像'" json:"avatar"`
	Gender    string         `gorm:"comment:'性别'" json:"gender"`
	Level     uint8          `gorm:"comment:'等级'" json:"level"`
	JoinTime  time.Time      `gorm:"comment:'加入时间'" json:"join_time"`
	Silence   bool           `gorm:"comment:'禁言'" json:"silence"`
}

func (u User) CheckPass(password string) bool {
	if u.Password == password {
		return true
	}
	return false
}

func (u User) GetUserID() uint64 {
	return u.UID
}
