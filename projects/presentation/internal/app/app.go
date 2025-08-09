package app

import (
	"hamroh/presentation/internal/configs"
	"hamroh/presentation/internal/handlers/http"
	"hamroh/presentation/internal/repositories"
	"hamroh/presentation/internal/services"
	"hamroh/presentation/internal/use-cases/presentation"

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

	presentationRepository := repositories.New(app.db)
	presentationService := services.New(presentationRepository)
	presentationUseCase := presentation.New(presentationService)
	presentationHandler := http.New(presentationUseCase)

	app.router = app.LoadRoutes(presentationHandler)

	return app
}
