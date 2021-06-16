package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Account 账号信息
type Account struct {
	UID       uint64    `db:"uid"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	Verified  bool      `db:"verified"`
	State     uint8     `db:"state"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// SetPassword 设置密码
func (acc *Account) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), EncryptCost)
	if err != nil {
		return err
	}
	acc.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (acc *Account) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	return err == nil
}

// BindAccount 使用UID绑定数据到结构体
func (acc *Account) BindAccount(uid interface{}) (err error) {
	sqlStr := "SELECT * FROM accounts WHERE uid = ?"
	err = DB.Get(acc, sqlStr, uid)
	return
}

// BindAccountByUsername 使用UID绑定数据到结构体
func (acc *Account) BindAccountByUsername(username string) (err error) {
	sqlStr := "SELECT * FROM accounts WHERE username = ?"
	err = DB.Get(acc, sqlStr, username)
	return
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *Account {
	if acc, _ := c.Get("USER"); acc != nil {
		if a, ok := acc.(*Account); ok {
			return a
		}
	}
	return nil
}
