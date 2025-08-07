package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kodnavis/people/internal/configs"
	"kodnavis/people/internal/handlers/http"
	"kodnavis/people/internal/repositories"
	"kodnavis/people/internal/services"
	"kodnavis/people/internal/use-cases/people"
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

	peopleRepository := repositories.New(app.db)
	peopleService := services.New(peopleRepository)
	peopleUseCase := people.New(peopleService)
	peopleHandler := http.New(peopleUseCase)

	app.router = app.LoadRoutes(peopleHandler)

	return app
}
