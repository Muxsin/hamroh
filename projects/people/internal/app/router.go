package app

import "github.com/gin-gonic/gin"

func (a *app) LoadRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/people", func(c *gin.Context) {
		c.JSON(200, a.configs.AppName)
	})

	return router
}
