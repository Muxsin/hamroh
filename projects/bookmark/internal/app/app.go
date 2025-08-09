package app

import (
	"hamroh/bookmark/internal/configs"
	"hamroh/bookmark/internal/handlers/http"
	"hamroh/bookmark/internal/repositories"
	"hamroh/bookmark/internal/services"
	"hamroh/bookmark/internal/use-cases/bookmark"

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

	bookmarkRepository := repositories.New(app.db)
	bookmarkService := services.New(bookmarkRepository)
	bookmarkUseCase := bookmark.New(bookmarkService)
	bookmarkHandler := http.New(bookmarkUseCase)

	app.router = app.LoadRoutes(bookmarkHandler)

	return app
}
