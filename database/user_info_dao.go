package database

import (
	"database/sql"

	"acgfate/model"
)

type UserInfoDao struct{}

func (d *UserInfoDao) QueryRow(idx interface{}) (interface{}, error) {
	sqlStr := "SELECT * FROM af_user WHERE uid = ?"
	ret := new(model.UserInfo)
	err := db.Get(ret, sqlStr, idx)
	return ret, err
}

func (d *UserInfoDao) InsertRow(tx *sql.Tx, uid int64) error {
	sqlStr := "INSERT INTO af_user_info(uid) VALUES (?)"
	_, err := tx.Exec(sqlStr, uid)
	return err
}

func (d *UserInfoDao) UpdateRow() {}

func (d *UserInfoDao) DeleteRow() {}

func (d *UserInfoDao) DeleteMRow() {}
