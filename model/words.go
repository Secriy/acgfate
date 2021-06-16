package model

import (
	"time"
)

// Words 文字信息
type Words struct {
	WID       uint64    `db:"wid"`
	UID       uint64    `db:"uid"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	Reports   uint      `db:"reports"`
	Views     uint      `db:"views"`
	Likes     uint      `db:"likes"`
	Dislikes  uint      `db:"dislikes"`
	Comments  uint      `db:"comments"`
	Category  uint8     `db:"category"`
	Tags      string    `db:"tags"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

// WidGet 使用WID获取Words模型
func (w *Words) WidGet(wid interface{}) (err error) {
	sqlStr := "SELECT * from words where wid = ?"
	err = DB.Get(w, sqlStr, wid)
	return
}
