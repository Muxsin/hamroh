package main

import (
	"hamroh/grocery/internal/database"
	"hamroh/grocery/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Grocery{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
