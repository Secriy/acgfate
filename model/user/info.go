package user

import (
	"time"

	"acgfate/model"
	"github.com/gin-gonic/gin"
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
	StatusNormal = iota // 状态: 正常
	StatusBanned        // 状态: 被封禁
)

// 状态对照表
var _statusFlags = map[uint8]string{
	StatusNormal: "账号正常",
	StatusBanned: "账号被封禁",
}

// 等级对照表
var _userLevelMap = []int{
	0,    // 0
	10,   // 1
	220,  // 2
	450,  // 3
	890,  // 4
	1500, // 5
	2900, // 6
	4700, // 7
	7800, // 8
	9600, // 9
}

// BaseInfo 用户信息结构
type BaseInfo struct {
	UID           uint64    `db:"uid"`
	Username      string    `db:"username"`
	Password      string    `db:"password"`
	Nickname      string    `db:"nickname"`
	Email         string    `db:"email"`
	EmailVerified bool      `db:"email_verified"`
	JoinTime      time.Time `db:"join_time"`
	AccountState  uint8     `db:"account_state"`
	Exp           int       `db:"exp"`
	Sign          string    `db:"sign"`
	Gender        uint8     `db:"gender"`
	Credit        uint32    `db:"credit"`
	Birthday      time.Time `db:"birthday"`
	Province      string    `db:"province"`
	City          string    `db:"city"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// GetBaseInfo 使用UID获取用户基本信息模型
func (info *BaseInfo) GetBaseInfo(c *gin.Context) (err error) {
	uid := c.GetUint64("UID")
	sqlStr := "SELECT * from user_base_info where uid = ?"
	err = model.DB.Get(info, sqlStr, uid)
	return
}

// GetState 使用状态信息查找对应状态
func GetState(num uint8) (status string) {
	status = _statusFlags[num]
	return
}

// GetUserInfoByID 获取当前用户信息模型
func GetUserInfoByID(uid interface{}) (userInfo BaseInfo, err error) {
	sqlStr := "SELECT * from user_base_info where uid = ?"
	err = model.DB.Get(&userInfo, sqlStr, uid)
	return
}

// GetLevelByExp 根据经验值获取用户等级
func GetLevelByExp(exp int) (level int) {
	for rLevel, v := range _userLevelMap {
		if exp < v {
			level = rLevel - 1
			break
		}
	}
	return
}
