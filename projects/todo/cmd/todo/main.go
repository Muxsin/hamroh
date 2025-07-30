package main

import (
	"hamroh-todo/internal/app"
	"hamroh-todo/internal/configs"
)

func main() {
	appConfigs := configs.New()
	myApp := app.New(appConfigs)

	if err := myApp.Run(); err != nil {
		panic(err)
	}
}
