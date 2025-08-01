package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, "Todo deleted!")
}
