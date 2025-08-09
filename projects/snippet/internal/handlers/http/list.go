package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type snippetResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	snippets, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []snippetResponse

	for _, snippet := range snippets {
		response = append(response, snippetResponse{
			Id:        snippet.ID,
			Text:      snippet.Text,
			CreatedAt: snippet.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
