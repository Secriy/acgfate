package database

import (
	"database/sql"

	"acgfate/model"
	"go.uber.org/zap"
)

type CateDao struct{}

// QueryByID use CID to query a row
func (d *CateDao) QueryByID(idx interface{}) (ret *model.Category, err error) {
	sqlStr := "SELECT * FROM af_category WHERE cate_id = ?"
	ret = new(model.Category)
	err = db.Get(ret, sqlStr, idx)
	return
}

// QueryByCname use CID to query a row
func (d *CateDao) QueryByCname(cname string) (ret *model.Category, err error) {
	sqlStr := "SELECT * FROM af_category WHERE cate_name  = ?"
	ret = new(model.Category)
	err = db.Get(ret, sqlStr, cname)
	return
}

// QueryAll get all the categories
func (d *CateDao) QueryAll() (ret []*model.Category, err error) {
	sqlStr := "SELECT * FROM  af_category"
	err = db.Select(&ret, sqlStr)
	if err == sql.ErrNoRows {
		zap.S().Warn("category table empty")
		err = nil
	}
	return
}
