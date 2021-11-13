package database

import (
	"database/sql"

	"acgfate/model"
)

type UserInfoDao struct{}

func (d *UserInfoDao) Query(idx interface{}) (*model.UserInfo, error) {
	sqlStr := "SELECT * FROM af_user WHERE uid = ?"
	ret := new(model.UserInfo)
	err := db.Get(ret, sqlStr, idx)
	return ret, err
}

func (d *UserInfoDao) Insert(tx *sql.Tx, uid int64) error {
	sqlStr := "INSERT INTO af_user_info(uid) VALUES (?)"
	_, err := tx.Exec(sqlStr, uid)
	return err
}

func (d *UserInfoDao) Update() {}

func (d *UserInfoDao) Delete() {}

func (d *UserInfoDao) MDelete() {}
