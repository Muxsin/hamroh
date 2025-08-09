package http

import (
	"hamroh/palette/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPaletteRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createPaletteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	palette := &models.Palette{
		Text: req.Text,
	}

	if err := h.useCase.Create(palette); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, palette)
}
