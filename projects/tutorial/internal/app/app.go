package app

import (
	"hamroh/tutorial/internal/configs"
	"hamroh/tutorial/internal/handlers/http"
	"hamroh/tutorial/internal/repositories"
	"hamroh/tutorial/internal/services"
	"hamroh/tutorial/internal/use-cases/tutorial"

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

	tutorialRepository := repositories.New(app.db)
	tutorialService := services.New(tutorialRepository)
	tutorialUseCase := tutorial.New(tutorialService)
	tutorialHandler := http.New(tutorialUseCase)

	app.router = app.LoadRoutes(tutorialHandler)

	return app
}
