package database

import (
	"database/sql"

	"acgfate/model"
)

type CatDao struct{}

// QueryByID use CID to query a row
func (d *CatDao) QueryByID(idx interface{}) (ret *model.Category, err error) {
	sqlStr := "SELECT * FROM af_category WHERE cat_id = ?"
	ret = new(model.Category)
	err = db.Get(ret, sqlStr, idx)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

// QueryByCname use CID to query a row
func (d *CatDao) QueryByCname(cname string) (ret *model.Category, err error) {
	sqlStr := "SELECT * FROM af_category WHERE cat_name  = ?"
	ret = new(model.Category)
	err = db.Get(ret, sqlStr, cname)
	return
}

// QueryAll get all the categories
func (d *CatDao) QueryAll() (ret []*model.Category, err error) {
	sqlStr := "SELECT * FROM  af_category"
	err = db.Select(&ret, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}
