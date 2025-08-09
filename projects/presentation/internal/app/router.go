package app

import (
	"github.com/gin-gonic/gin"
)

type presentationHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler presentationHandler) *gin.Engine {
	router := gin.New()

	router.GET("/presentation", handler.List)
	router.POST("/presentation", handler.Create)
	router.GET("/presentation/:id", handler.Get)
	router.DELETE("/presentation/:id", handler.Delete)
	router.PUT("/presentation/:id", handler.Update)

	return router
}
