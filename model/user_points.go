package model

import "time"

const (
	userLevelMax = 100
)

var userLevelMap = []int64{
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

type UserPoints struct {
	UID      uint64    `gorm:"primaryKey;unique;autoIncrement;comment:'用户ID'"`
	EXP      int64     `gorm:"comment:'经验值';default:0"`
	Coins    uint      `gorm:"comment:'点数';default:0"`
	SignTime time.Time `gorm:"comment:'签到时间'"`
}

// FormatLevel 获取用户等级
func FormatLevel(exp int64) (level uint8) {
	for rLevel, v := range userLevelMap {
		if exp < v {
			level = uint8(rLevel) - 1
			break
		}
	}
	return
}
