package http

import (
	"hamroh/presentation/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPresentationRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createPresentationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	presentation := &models.Presentation{
		Text: req.Text,
	}

	if err := h.useCase.Create(presentation); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, presentation)
}
