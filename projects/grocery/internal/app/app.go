package app

import (
	"hamroh/grocery/internal/configs"
	"hamroh/grocery/internal/handlers/http"
	"hamroh/grocery/internal/repositories"
	"hamroh/grocery/internal/services"
	"hamroh/grocery/internal/use-cases/grocery"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type app struct {
	configs *configs.Configs
	router  *gin.Engine
	db      *gorm.DB
}

func New(configs *configs.Configs, db *gorm.DB) *app {
	app := &app{
		configs: configs,
		db:      db,
	}

	groceryRepository := repositories.New(app.db)
	groceryService := services.New(groceryRepository)
	groceryUseCase := grocery.New(groceryService)
	groceryHandler := http.New(groceryUseCase)

	app.router = app.LoadRoutes(groceryHandler)

	return app
}
