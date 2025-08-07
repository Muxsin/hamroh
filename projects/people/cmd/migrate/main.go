package main

import (
	"github.com/joho/godotenv"
	"kodnavis/people/internal/database"
	"kodnavis/people/internal/models"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.People{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
