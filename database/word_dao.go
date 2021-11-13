package database

import (
	"database/sql"

	"acgfate/model"
)

type WordDao struct{}

// QueryByID query single row by ID.
func (w *WordDao) QueryByID(wid interface{}) (ret *model.Word, err error) {
	ret = new(model.Word)
	sqlStr := "SELECT * FROM  af_word WHERE wid = ?"
	err = db.Get(ret, sqlStr, wid)
	return
}

// MQuery query multi rows by limit and offset.
func (w *WordDao) MQuery(offset, limit int64) (ret []*model.Word, err error) {
	sqlStr := "SELECT * FROM  af_word LIMIT ?,?"
	err = db.Select(&ret, sqlStr, offset*limit, limit)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

// MQueryByAuthor query multi rows by author ID.
func (w *WordDao) MQueryByAuthor(uid interface{}, offset, limit int64) (ret []*model.Word, err error) {
	sqlStr := "SELECT * FROM  af_word WHERE aid = ? LIMIT ?,?"
	err = db.Select(&ret, sqlStr, uid, offset*limit, limit)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

// MQueryByCat query multi rows by category ID.
func (w *WordDao) MQueryByCat(catID, offset, limit int64) (ret []*model.Word, err error) {
	sqlStr := "SELECT * FROM  af_word WHERE cat_id = ? LIMIT ?,?"
	err = db.Select(&ret, sqlStr, catID, offset*limit, limit)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func (w *WordDao) Insert(word *model.Word) error {
	sqlStr := "INSERT INTO af_word(wid, aid, cat_id, title, content) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, word.Wid, word.Aid, word.CatID, word.Title, word.Content)
	return err
}
