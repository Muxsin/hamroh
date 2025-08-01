package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "Todo updated!")
}
