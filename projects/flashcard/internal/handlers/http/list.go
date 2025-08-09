package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type flashcardResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	flashcards, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []flashcardResponse

	for _, flashcard := range flashcards {
		response = append(response, flashcardResponse{
			Id:        flashcard.ID,
			Text:      flashcard.Text,
			CreatedAt: flashcard.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
