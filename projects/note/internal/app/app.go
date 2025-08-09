package app

import (
	"hamroh/note/internal/configs"
	"hamroh/note/internal/handlers/http"
	"hamroh/note/internal/repositories"
	"hamroh/note/internal/services"
	"hamroh/note/internal/use-cases/note"

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

	noteRepository := repositories.New(app.db)
	noteService := services.New(noteRepository)
	noteUseCase := note.New(noteService)
	noteHandler := http.New(noteUseCase)

	app.router = app.LoadRoutes(noteHandler)

	return app
}
