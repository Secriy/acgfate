package model

import (
	"fmt"

	"acgfate/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	dsn := config.Conf.DSN
	mysqlDB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Disabling MySQL tests:\n    %v", err))
	}
	mysqlDB.SetMaxOpenConns(200)
	mysqlDB.SetMaxIdleConns(10)

	DB = mysqlDB
}
