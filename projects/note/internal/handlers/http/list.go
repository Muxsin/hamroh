package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type noteResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	notes, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []noteResponse

	for _, note := range notes {
		response = append(response, noteResponse{
			Id:        note.ID,
			Text:      note.Text,
			CreatedAt: note.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
