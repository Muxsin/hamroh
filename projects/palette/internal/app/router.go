package app

import (
	"github.com/gin-gonic/gin"
)

type paletteHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler paletteHandler) *gin.Engine {
	router := gin.New()

	router.GET("/palette", handler.List)
	router.POST("/palette", handler.Create)
	router.GET("/palette/:id", handler.Get)
	router.DELETE("/palette/:id", handler.Delete)
	router.PUT("/palette/:id", handler.Update)

	return router
}
