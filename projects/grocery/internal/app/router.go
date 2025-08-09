package app

import (
	"github.com/gin-gonic/gin"
)

type groceryHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler groceryHandler) *gin.Engine {
	router := gin.New()

	router.GET("/grocery", handler.List)
	router.POST("/grocery", handler.Create)
	router.GET("/grocery/:id", handler.Get)
	router.DELETE("/grocery/:id", handler.Delete)
	router.PUT("/grocery/:id", handler.Update)

	return router
}
