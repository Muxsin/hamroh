package app

import (
	"github.com/gin-gonic/gin"
)

type peopleHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler peopleHandler) *gin.Engine {
	router := gin.New()

	router.GET("/people", handler.List)
	router.POST("/people", handler.Create)
	router.GET("/people/:id", handler.Get)
	router.DELETE("/people/:id", handler.Delete)
	router.PUT("/people/:id", handler.Update)

	return router
}
