package model

import (
	"database/sql"
	"time"
)

const (
	userLevelMax = 100
)

var userLevelMap = []uint16{
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
	UID       uint64       `db:"uid"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	IsDeleted bool         `db:"is_deleted"`
	EXP       uint16       `db:"exp"`
	Coins     uint64       `db:"coins"`
	SignTime  sql.NullTime `db:"sign_time"`
}

// GetUserPoint 获取当前用户点数模型
func GetUserPoint(uid interface{}) (userPoints UserPoints, err error) {
	infoStr := "SELECT * from user_point where uid = ?"
	err = DB.Get(&userPoints, infoStr, uid)
	return
}

// FormatLevel 获取用户等级
func FormatLevel(exp uint16) (level uint8) {
	for rLevel, v := range userLevelMap {
		if exp < v {
			level = uint8(rLevel) - 1
			break
		}
	}
	return
}
