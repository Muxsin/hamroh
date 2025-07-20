package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database")

	return db
}
