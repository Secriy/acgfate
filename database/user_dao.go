package database

import (
	"acgfate/model"
)

type UserDao struct{}

// QueryByUID query single row by UID
func (d *UserDao) QueryByUID(idx interface{}) (ret *model.User, err error) {
	ret = new(model.User)
	sqlStr := "SELECT * FROM af_user WHERE uid = ?"
	err = db.Get(ret, sqlStr, idx)
	return
}

// QueryByUname query single row by username
func (d *UserDao) QueryByUname(idx interface{}) (ret *model.User, err error) {
	ret = new(model.User)
	sqlStr := "SELECT * FROM af_user WHERE username = ?"
	err = db.Get(ret, sqlStr, idx)
	return
}

// QueryByEmail query single row by email
func (d *UserDao) QueryByEmail(idx interface{}) (ret *model.User, err error) {
	ret = new(model.User)
	sqlStr := "SELECT * FROM af_user WHERE email = ?"
	err = db.Get(ret, sqlStr, idx)
	return
}

func (d *UserDao) Insert(user *model.User) error {
	tx, err := db.Begin()
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
	if err := infoDao.Insert(tx, uid); err != nil {
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()
	return nil
}

func (d *UserDao) Update() {}

func (d *UserDao) Delete() {}

func (d *UserDao) MDelete() {}

// NickName 获取用户的昵称
func NickName(uid interface{}) string {
	dao := new(UserDao)
	user, err := dao.QueryByUID(uid)
	if err != nil {
		return ""
	}
	return user.Nickname
}
