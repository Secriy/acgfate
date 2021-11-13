package model

import "time"

// Category 分区模型
type Category struct {
	ID           int64     `db:"id"`
	CategoryID   int64     `db:"cat_id"`
	CategoryName string    `db:"cat_name"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
