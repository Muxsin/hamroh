package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, "Todo updated!")
}
