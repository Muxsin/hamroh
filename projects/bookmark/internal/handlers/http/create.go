package http

import (
	"hamroh/bookmark/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createBookmarkRequest struct {
	Text string `json:"text" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createBookmarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookmark := &models.Bookmark{
		Text: req.Text,
	}

	if err := h.useCase.Create(bookmark); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, bookmark)
}
