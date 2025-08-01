package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hamroh-todo/internal/handlers/http"
	"hamroh-todo/internal/repositories"
	"hamroh-todo/internal/services"
	use_cases "hamroh-todo/internal/use-cases"
)

type configs interface {
}

type App struct {
	configs configs
	router  *gin.Engine
	db      *gorm.DB
}

func New(configs configs, db *gorm.DB) *App {
	app := &App{
		configs: configs,
		router:  gin.New(),
		db:      db,
	}

	todoRepository := repositories.New(app.db)
	todoService := services.New(todoRepository)
	todoUseCase := use_cases.New(todoService)
	todoHandler := http.New(todoUseCase)

	app.router = app.LoadRoutes(todoHandler)

	return app
}
