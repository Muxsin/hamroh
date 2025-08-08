package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type peopleResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birthdate time.Time `json:"birthdate"`
	CreatedAt string    `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	peoples, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []peopleResponse

	for _, people := range peoples {
		response = append(response, peopleResponse{
			Id:        people.ID,
			Name:      people.Name,
			Surname:   people.Surname,
			Birthdate: people.Birthdate,
			CreatedAt: people.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
