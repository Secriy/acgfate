package database

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
