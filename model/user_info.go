package model

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	EncryptCost = 12 // 加密强度
)

const (
	Private = iota // 性别: 保密
	Male           // 性别: 男
	Female         // 性别: 女
	Other          // 性别: 其他
)

const (
	StatusCommon   = iota // 状态: 正常
	StatusSilenced        // 状态: 被禁言
	StatusBanned          // 状态: 被封禁
)

var statusFlags = map[uint8]string{
	StatusCommon:   "账号正常",
	StatusSilenced: "账号被禁言",
	StatusBanned:   "账号被封禁",
}

var genderFlags = map[uint8]string{
	Private: "保密",
	Male:    "男",
	Female:  "女",
	Other:   "其他",
}

type UserInfo struct {
	UID          uint64         `db:"uid"`
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
	IsDeleted    bool           `db:"is_deleted"`
	Username     string         `db:"username"`
	Password     string         `db:"password"`
	Nickname     string         `db:"nickname"`
	Mail         string         `db:"mail"`
	MailVerified bool           `db:"mail_verify"`
	Status       uint8          `db:"status"`
	Avatar       string         `db:"avatar"`
	Sign         sql.NullString `db:"sign"`
	Gender       uint8          `db:"gender"`
	Birthday     sql.NullString `db:"birthday"`
}

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

// GetUserInfo 获取当前用户信息模型
func GetUserInfo(uid interface{}) (userInfo UserInfo, err error) {
	sqlStr := "SELECT * from user_info where uid = ?"
	err = DB.Get(&userInfo, sqlStr, uid)
	return
}

// GetStatus 使用状态信息查找对应状态
func GetStatus(statusNum uint8) (status string) {
	status = statusFlags[statusNum]
	return
}

// GetGender 查询对应性别
func GetGender(genderNum uint8) (gender string) {
	gender = genderFlags[genderNum]
	return
}
