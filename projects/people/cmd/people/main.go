package main

import (
	"github.com/joho/godotenv"
	"kodnavis/people/internal/app"
	"kodnavis/people/internal/configs"
	"kodnavis/people/internal/database"
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
