package model

type Words struct {
	WID       uint64 `db:"wid"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	IsDeleted bool   `db:"is_deleted"`
	Publisher uint64 `db:"publisher"`
	Content   string `db:"content"`
}

// GetWordsByWID 根据文章ID获取文章
func GetWordsByWID(uid interface{}) (words Words, err error) {
	infoStr := "SELECT * from words where wid = ?"
	err = DB.Get(&words, infoStr, uid)
	return
}
