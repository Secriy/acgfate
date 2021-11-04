package database

import (
	"errors"

	"acgfate/model"
)

type UserDao struct{}

const (
	QUID = iota
	QUname
	QEmail
)

func (d *UserDao) QueryRow(col int, idx interface{}) (*model.User, error) {
	var user model.User
	var sqlStr string
	switch col {
	case QUname:
		sqlStr = "SELECT * FROM af_user WHERE username = ?"
	case QEmail:
		sqlStr = "SELECT * FROM af_user WHERE email = ?"
	default:
		sqlStr = "SELECT * FROM af_user WHERE uid = ?"
	}
	err := DB.Get(&user, sqlStr, idx)
	return &user, err
}

func (d *UserDao) InsertRow(userRow interface{}) error {
	if user, ok := userRow.(model.User); ok {
		tx, err := DB.Begin()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		sqlStr := "INSERT INTO af_user(username, password, nickname, email) VALUES (?, ?, ?, ?)"
		ret, err := tx.Exec(sqlStr, user.Username, user.Password, user.Nickname, user.Email)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// get last insert row id
		uid, err := ret.LastInsertId()
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		// insert into info table
		infoDao := UserInfoDao{}
		if err := infoDao.InsertRow(tx, uid); err != nil {
			_ = tx.Rollback()
			return err
		}
		_ = tx.Commit()
		return nil
	}
	return errors.New("") // TODO: 错误信息
}

func (d *UserDao) UpdateRow() {}

func (d *UserDao) DeleteRow() {}

func (d *UserDao) DeleteMRow() {}

// IsExists 判断是否存在字段
func (d *UserDao) IsExists(col int, idx interface{}) bool {
	_, err := d.QueryRow(col, idx)
	return err == nil
}
