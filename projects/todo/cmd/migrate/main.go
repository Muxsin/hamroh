package main

import (
	"github.com/joho/godotenv"
	"hamroh-todo/internal/database"
	"hamroh-todo/internal/models"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
