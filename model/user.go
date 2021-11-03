package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// EncryptCost 加密难度
const EncryptCost = 12

// User 用户账号信息
type User struct {
	UID       uint64    `db:"uid"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Nickname  string    `db:"nickname"`
	Phone     string    `db:"phone"`
	Email     string    `db:"email"`
	Avatar    string    `db:"avatar"`
	State     uint8     `db:"state"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UserInfo 用户个人信息
type UserInfo struct {
	UID       uint64    `db:"uid"`
	Gender    uint8     `db:"gender"`
	Sign      string    `db:"sign"`
	Birthday  time.Time `db:"birthday"`
	Province  string    `db:"province"`
	City      string    `db:"city"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), EncryptCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
