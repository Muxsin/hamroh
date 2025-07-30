package app

import (
	"github.com/gin-gonic/gin"
	"hamroh-todo/internal/configs"
)

type App struct {
	configs *configs.Configs
	router  *gin.Engine
}

func New(configs *configs.Configs) *App {
	app := &App{
		configs: configs,
	}

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Done!")
	})

	app.router = router

	return app
}
