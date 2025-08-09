package main

import (
	"kodnavis/password/internal/database"
	"kodnavis/password/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Password{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
