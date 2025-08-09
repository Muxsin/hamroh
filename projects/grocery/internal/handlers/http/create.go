package http

import (
	"hamroh/grocery/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createGroceryRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createGroceryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grocery := &models.Grocery{
		Text: req.Text,
	}

	if err := h.useCase.Create(grocery); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, grocery)
}
