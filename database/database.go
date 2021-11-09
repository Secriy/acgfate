package database

import (
	"fmt"

	"acgfate/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitMySQL 初始化数据库连接
func InitMySQL(conf *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Passwd,
		conf.Host,
		conf.Port,
		conf.DB,
	)
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)

	return
}

func CloseMySQL() {
	_ = db.Close()
}
