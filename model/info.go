package model

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	EncryptCost = 12 // 加密强度
)

// BasicInfo 用户信息结构
type BasicInfo struct {
	UID       uint64    `db:"uid"`
	Nickname  string    `db:"nickname"`
	JoinTime  time.Time `db:"join_time"`
	Exp       int       `db:"exp"`
	Sign      string    `db:"sign"`
	Gender    uint8     `db:"gender"`
	Credit    uint32    `db:"credit"`
	Birthday  time.Time `db:"birthday"`
	Province  string    `db:"province"`
	City      string    `db:"city"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// CurrentBasicInfo 使用UID获取用户基本信息模型
func (info *BasicInfo) CurrentBasicInfo(c *gin.Context) (username string, err error) {
	acc := CurrentUser(c)
	username = acc.Username
	sqlStr := "SELECT * from user_basic_info where uid = ?"
	err = DB.Get(info, sqlStr, acc.UID)
	return
}

// GetState 使用状态信息查找对应状态
func GetState(num uint8) (status string) {
	status = _statusFlags[num]
	return
}

// GetUserInfoByID 获取当前用户信息模型
func GetUserInfoByID(uid interface{}) (userInfo BasicInfo, err error) {
	sqlStr := "SELECT * from user_basic_info where uid = ?"
	err = DB.Get(&userInfo, sqlStr, uid)
	return
}
