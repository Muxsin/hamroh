package main

import (
	"hamroh/note/internal/database"
	"hamroh/note/internal/models"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	db := database.Connect()

	if err := db.AutoMigrate(&models.Note{}); err != nil {
		panic(err)
	}

	log.Println("migrate the schemas finished")
}
