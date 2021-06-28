package model

import (
	"context"
	"fmt"

	"acgfate/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	InitMongoDB()
}

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS!")
}
