package main

import (
	"hamroh/presentation/internal/database"
	"hamroh/presentation/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Presentation{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
