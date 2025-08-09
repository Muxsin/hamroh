package main

import (
	"hamroh/snippet/internal/database"
	"hamroh/snippet/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Snippet{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
