package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type passwordResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	passwords, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []passwordResponse

	for _, password := range passwords {
		response = append(response, passwordResponse{
			Id:        password.ID,
			Text:      password.Text,
			CreatedAt: password.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
