package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect to mysql

var DB *gorm.DB

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/socialApp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Connected to DB")
}

func GetDB() *gorm.DB {
	return DB
}
