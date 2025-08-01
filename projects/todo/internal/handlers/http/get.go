package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, "Todo gotten!")
}
