package main

import (
	"hamroh/palette/internal/database"
	"hamroh/palette/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Palette{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
