package app

import (
	"github.com/gin-gonic/gin"
)

type recipeHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler recipeHandler) *gin.Engine {
	router := gin.New()

	router.GET("/recipe", handler.List)
	router.POST("/recipe", handler.Create)
	router.GET("/recipe/:id", handler.Get)
	router.DELETE("/recipe/:id", handler.Delete)
	router.PUT("/recipe/:id", handler.Update)

	return router
}
