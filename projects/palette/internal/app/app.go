package app

import (
	"hamroh/palette/internal/configs"
	"hamroh/palette/internal/handlers/http"
	"hamroh/palette/internal/repositories"
	"hamroh/palette/internal/services"
	"hamroh/palette/internal/use-cases/palette"

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

	paletteRepository := repositories.New(app.db)
	paletteService := services.New(paletteRepository)
	paletteUseCase := palette.New(paletteService)
	paletteHandler := http.New(paletteUseCase)

	app.router = app.LoadRoutes(paletteHandler)

	return app
}
