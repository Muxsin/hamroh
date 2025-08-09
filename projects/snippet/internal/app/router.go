package app

import (
	"github.com/gin-gonic/gin"
)

type snippetHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler snippetHandler) *gin.Engine {
	router := gin.New()

	router.GET("/snippet", handler.List)
	router.POST("/snippet", handler.Create)
	router.GET("/snippet/:id", handler.Get)
	router.DELETE("/snippet/:id", handler.Delete)
	router.PUT("/snippet/:id", handler.Update)

	return router
}
