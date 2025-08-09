package app

import (
	"hamroh/snippet/internal/configs"
	"hamroh/snippet/internal/handlers/http"
	"hamroh/snippet/internal/repositories"
	"hamroh/snippet/internal/services"
	"hamroh/snippet/internal/use-cases/snippet"

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

	snippetRepository := repositories.New(app.db)
	snippetService := services.New(snippetRepository)
	snippetUseCase := snippet.New(snippetService)
	snippetHandler := http.New(snippetUseCase)

	app.router = app.LoadRoutes(snippetHandler)

	return app
}
