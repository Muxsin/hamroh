package app

import (
	"github.com/gin-gonic/gin"
)

type tutorialHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler tutorialHandler) *gin.Engine {
	router := gin.New()

	router.GET("/tutorial", handler.List)
	router.POST("/tutorial", handler.Create)
	router.GET("/tutorial/:id", handler.Get)
	router.DELETE("/tutorial/:id", handler.Delete)
	router.PUT("/tutorial/:id", handler.Update)

	return router
}
