package database

import (
	"database/sql"

	"acgfate/model"
)

type WordDao struct{}

// QueryByID query single row by ID.
func (w *WordDao) QueryByID(wid interface{}) (ret *model.Word, err error) {
	ret = new(model.Word)
	sqlStr := `SELECT * FROM  af_word WHERE wid = ?`
	err = db.Get(ret, sqlStr, wid)
	if err == sql.ErrNoRows {
		ret, err = nil, nil
	}
	return
}

// MQuery query multi rows by limit and offset.
func (w *WordDao) MQuery(offset, limit int64) (ret []*model.Word, err error) {
	sqlStr := `SELECT * FROM  af_word LIMIT ?,?`
	err = db.Select(&ret, sqlStr, offset*limit, limit)
	if err == sql.ErrNoRows {
		ret, err = nil, nil
	}
	return
}

// MQueryByAuthor query multi rows by author ID.
func (w *WordDao) MQueryByAuthor(uid interface{}, offset, limit int64) (ret []*model.Word, err error) {
	sqlStr := `SELECT * FROM  af_word WHERE aid = ? LIMIT ?,?`
	err = db.Select(&ret, sqlStr, uid, offset*limit, limit)
	if err == sql.ErrNoRows {
		ret, err = nil, nil
	}
	return
}

// MQueryByCat query multi rows by category ID.
func (w *WordDao) MQueryByCat(catID, offset, limit int64) (ret []*model.Word, err error) {
	sqlStr := `SELECT * FROM  af_word WHERE cat_id = ? LIMIT ?,?`
	err = db.Select(&ret, sqlStr, catID, offset*limit, limit)
	if err == sql.ErrNoRows {
		ret, err = nil, nil
	}
	return
}

func (w *WordDao) Insert(word *model.Word) (err error) {
	sqlStr := `INSERT INTO af_word(wid, aid, cat_id, title, content) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, word.Wid, word.Aid, word.CatID, word.Title, word.Content)
	return
}

func (w *WordDao) Update(word *model.Word) (err error) {
	sqlStr := `UPDATE af_word SET aid =?, cat_id =?, status =?, title =?, content =? WHERE wid =?`
	_, err = db.Exec(
		sqlStr,
		word.Aid,
		word.CatID,
		word.Status,
		word.Title,
		word.Content,
		word.Wid,
	)
	return
}
