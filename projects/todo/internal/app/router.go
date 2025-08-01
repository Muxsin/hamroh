package app

import (
	"github.com/gin-gonic/gin"
)

type handler interface {
	Create(*gin.Context)
	List(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func (a *App) LoadRoutes(h handler) *gin.Engine {
	router := gin.New()
	router.GET("/todos", h.List)
	router.GET("/todos/:id", h.Get)
	router.POST("/todos", h.Create)
	router.PUT("/todos/:id", h.Update)

	return router
}
