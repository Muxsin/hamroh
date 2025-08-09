package http

import (
	"hamroh/recipe/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRecipeRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createRecipeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := &models.Recipe{
		Text: req.Text,
	}

	if err := h.useCase.Create(recipe); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, recipe)
}
