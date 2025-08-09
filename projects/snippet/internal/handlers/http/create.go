package http

import (
	"hamroh/snippet/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createSnippetRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createSnippetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	snippet := &models.Snippet{
		Text: req.Text,
	}

	if err := h.useCase.Create(snippet); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, snippet)
}
