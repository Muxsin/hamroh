package http

import (
	"hamroh/tutorial/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createTutorialRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createTutorialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tutorial := &models.Tutorial{
		Text: req.Text,
	}

	if err := h.useCase.Create(tutorial); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, tutorial)
}
