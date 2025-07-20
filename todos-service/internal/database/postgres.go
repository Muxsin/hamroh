package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_POSTGRES")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")

	return db
}
