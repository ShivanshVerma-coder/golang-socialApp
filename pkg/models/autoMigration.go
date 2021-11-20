package models

import (
	"fmt"

	"github.com/ShivanshVerma-coder/golang-socialApp/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	fmt.Println("Auto Migration ran")
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&User{}, &Blog{}) // First Blog then User migrated
	if err != nil {
		panic(err)
	}
	// db.Model(&Blog{}).Exec("ALTER TABLE blogs ADD FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE")
}
