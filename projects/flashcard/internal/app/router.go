package app

import (
	"github.com/gin-gonic/gin"
)

type flashcardHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func (a *app) LoadRoutes(handler flashcardHandler) *gin.Engine {
	router := gin.New()

	router.GET("/flashcard", handler.List)
	router.POST("/flashcard", handler.Create)
	router.GET("/flashcard/:id", handler.Get)
	router.DELETE("/flashcard/:id", handler.Delete)
	router.PUT("/flashcard/:id", handler.Update)

	return router
}
