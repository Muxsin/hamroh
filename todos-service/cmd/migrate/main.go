package main

import (
	"github.com/joho/godotenv"
	"log"
	"todos-service/internal/database"
	"todos-service/internal/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		panic(err)
	}

	log.Println("Database migration succeeded")
}
