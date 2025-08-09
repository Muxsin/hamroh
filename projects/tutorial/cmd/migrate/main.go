package main

import (
	"hamroh/tutorial/internal/database"
	"hamroh/tutorial/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Tutorial{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
