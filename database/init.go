package database

import (
	"fmt"

	"acgfate/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// type DAO interface {
// 	// QueryRow 查询单条记录
// 	QueryRow(interface{}) (interface{}, error)
// 	// InsertRow 插入单条记录
// 	InsertRow(interface{}) (interface{}, error)
// 	// InsertMRow 插入多条记录
// 	InsertMRow([]interface{})
// 	// UpdateRow 更新单条记录
// 	UpdateRow()
// 	// DeleteRow 删除单条记录
// 	DeleteRow()
// 	// DeleteMRow 删除多条记录
// 	DeleteMRow()
// }

var DB *sqlx.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	dsn := config.Conf.DSN
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("connect to mysql service failed: %v\n", err))
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	DB = db
}

// func InitMongoDB() {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
//
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("SUCCESS!")
// }
