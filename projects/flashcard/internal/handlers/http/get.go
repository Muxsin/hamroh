package http

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getFlashcardResponse struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) Get(c *gin.Context) {
	id := c.Param("id")
	flashcard, err := h.useCase.GetOne(id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response := getFlashcardResponse{
		Id:        id,
		Text:      flashcard.Text,
		CreatedAt: flashcard.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
