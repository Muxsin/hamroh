package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN_POSTGRES")), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to database")

	return db, nil
}
