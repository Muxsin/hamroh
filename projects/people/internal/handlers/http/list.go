package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *handler) List(c *gin.Context) {
	peoples, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, peoples)
}
