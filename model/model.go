package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Error open database %v\n", err))
	}

	// 设置连接池
	// 空闲
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(50)
	// 打开
	sqlDB.SetMaxOpenConns(100)
	// 超时
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	_ = DB.AutoMigrate(&User{})
}
