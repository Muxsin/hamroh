package app

import (
	"github.com/gin-gonic/gin"
)

type bookmarkHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler bookmarkHandler) *gin.Engine {
	router := gin.New()

	router.GET("/bookmark", handler.List)
	router.POST("/bookmark", handler.Create)
	router.GET("/bookmark/:id", handler.Get)
	router.DELETE("/bookmark/:id", handler.Delete)
	router.PUT("/bookmark/:id", handler.Update)

	return router
}
