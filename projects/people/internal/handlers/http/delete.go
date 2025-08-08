package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.useCase.Delete(id); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
