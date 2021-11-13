package model

import "time"

// Category 用户账号信息
type Category struct {
	ID           int64     `db:"id"`
	CategoryID   int64     `db:"cate_id"`
	CategoryName string    `db:"cate_name"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
