package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
