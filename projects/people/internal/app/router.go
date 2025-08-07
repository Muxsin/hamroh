package app

import (
	"github.com/gin-gonic/gin"
)

type peopleHandler interface {
	Create(c *gin.Context)
}

func (a *app) LoadRoutes(handler peopleHandler) *gin.Engine {
	router := gin.New()

	router.GET("/people", func(c *gin.Context) {
		c.JSON(200, a.configs.AppName)
	})
	router.POST("/people", handler.Create)

	return router
}
