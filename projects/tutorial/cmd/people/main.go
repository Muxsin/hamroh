package main

import (
	"hamroh/tutorial/internal/app"
	"hamroh/tutorial/internal/configs"
	"hamroh/tutorial/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	config := configs.New()
	db := database.Connect()

	myApp := app.New(config, db)

	if err := myApp.Run(); err != nil {
		panic(err)
	}
}
