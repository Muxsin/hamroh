package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hamroh-todo/internal/app"
	"hamroh-todo/internal/configs"
)

func main() {
	appConfigs := configs.New()

	db, err := gorm.Open(postgres.Open(appConfigs.PostgresDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	myApp := app.New(appConfigs, db)

	if err := myApp.Run(); err != nil {
		panic(err)
	}
}
