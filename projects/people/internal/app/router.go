package app

import (
	"github.com/gin-gonic/gin"
)

type peopleHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
}

func (a *app) LoadRoutes(handler peopleHandler) *gin.Engine {
	router := gin.New()

	router.GET("/people", handler.List)
	router.POST("/people", handler.Create)

	return router
}
