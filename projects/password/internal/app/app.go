package app

import (
	"kodnavis/password/internal/configs"
	"kodnavis/password/internal/handlers/http"
	"kodnavis/password/internal/repositories"
	"kodnavis/password/internal/services"
	"kodnavis/password/internal/use-cases/password"

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

	passwordRepository := repositories.New(app.db)
	passwordService := services.New(passwordRepository)
	passwordUseCase := password.New(passwordService)
	passwordHandler := http.New(passwordUseCase)

	app.router = app.LoadRoutes(passwordHandler)

	return app
}
