package internal

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_NAME = "data.db"

	ERROR_DATABASE = "Internal Server Error 500"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	return db
}
