package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type presentationResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	presentations, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []presentationResponse

	for _, presentation := range presentations {
		response = append(response, presentationResponse{
			Id:        presentation.ID,
			Text:      presentation.Text,
			CreatedAt: presentation.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
