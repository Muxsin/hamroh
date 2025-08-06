package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kodnavis/people/internal/configs"
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

	app.router = app.LoadRoutes()

	return app
}
