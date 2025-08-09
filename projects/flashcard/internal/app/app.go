package app

import (
	"hamroh/flashcard/internal/configs"
	"hamroh/flashcard/internal/handlers/http"
	"hamroh/flashcard/internal/repositories"
	"hamroh/flashcard/internal/services"
	"hamroh/flashcard/internal/use-cases/flashcard"

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

	flashcardRepository := repositories.New(app.db)
	flashcardService := services.New(flashcardRepository)
	flashcardUseCase := flashcard.New(flashcardService)
	flashcardHandler := http.New(flashcardUseCase)

	app.router = app.LoadRoutes(flashcardHandler)

	return app
}
