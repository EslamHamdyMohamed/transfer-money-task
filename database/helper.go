package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"transfer_money/database/models/transfer"
)

var DB *gorm.DB
var err error

func ConnectToDatabase() {
	dsn := "root:1234567@tcp(127.0.0.1:3306)/transfer?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("ERROR", err.Error())
	} else {
		fmt.Println("Connected To Database")
	}
}

func MigrateAll() {
	DB.AutoMigrate(&transfer.Account{})
}
