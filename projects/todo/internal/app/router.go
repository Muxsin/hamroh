package app

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(*gin.Context)
	List(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func (a *App) LoadRoutes(h Handler) *gin.Engine {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, a.configs.AppName+" working!")
	})
	router.GET("/:id", h.Get)

	return router
}
