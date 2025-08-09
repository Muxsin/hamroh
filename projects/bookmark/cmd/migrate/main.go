package main

import (
	"hamroh/bookmark/internal/database"
	"hamroh/bookmark/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Bookmark{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
