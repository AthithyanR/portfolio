package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ChatMessage struct {
	Ip      string
	Name    string
	Message string
}

func InitDB() {
	dbConn, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = dbConn

	// DB.Create(&ChatMessages{
	// 	Ip:      "127.0.0.1",
	// 	Name:    "Athithyan",
	// 	Message: "Hello sqlite3",
	// })

	DB.AutoMigrate(&ChatMessage{})
}
