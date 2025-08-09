package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type groceryResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	grocerys, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []groceryResponse

	for _, grocery := range grocerys {
		response = append(response, groceryResponse{
			Id:        grocery.ID,
			Text:      grocery.Text,
			CreatedAt: grocery.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
