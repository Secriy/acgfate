package model

import "time"

// Category 用户账号信息
type Category struct {
	ID        int64     `db:"id"`
	CateID    int64     `db:"cate_id"`
	CateName  string    `db:"cate_name"`
	Desc      string    `db:"description"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
