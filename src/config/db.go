package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(DB_CONNECTION_STR), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
