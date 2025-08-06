package main

import (
	"hamroh-todo/internal/app"
	"hamroh-todo/internal/configs"
	"hamroh-todo/internal/database"
)

func main() {
	appConfigs := configs.New()
	db := database.Connect()

	if err := app.New(appConfigs, db).Run(); err != nil {
		panic(err)
	}
}
