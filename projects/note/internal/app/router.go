package app

import (
	"github.com/gin-gonic/gin"
)

type noteHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler noteHandler) *gin.Engine {
	router := gin.New()

	router.GET("/note", handler.List)
	router.POST("/note", handler.Create)
	router.GET("/note/:id", handler.Get)
	router.DELETE("/note/:id", handler.Delete)
	router.PUT("/note/:id", handler.Update)

	return router
}
