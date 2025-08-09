package app

import (
	"github.com/gin-gonic/gin"
)

type passwordHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler passwordHandler) *gin.Engine {
	router := gin.New()

	router.GET("/password", handler.List)
	router.POST("/password", handler.Create)
	router.GET("/password/:id", handler.Get)
	router.DELETE("/password/:id", handler.Delete)
	router.PUT("/password/:id", handler.Update)

	return router
}
