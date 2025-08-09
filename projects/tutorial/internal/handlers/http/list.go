package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type tutorialResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	tutorials, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []tutorialResponse

	for _, tutorial := range tutorials {
		response = append(response, tutorialResponse{
			Id:        tutorial.ID,
			Text:      tutorial.Text,
			CreatedAt: tutorial.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
