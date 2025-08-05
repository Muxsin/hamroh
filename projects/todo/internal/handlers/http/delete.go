package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.useCase.Delete(id); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
