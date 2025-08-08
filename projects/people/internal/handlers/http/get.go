package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type getPeopleResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birthdate time.Time `json:"birthdate"`
	CreatedAt string    `json:"created_at"`
}

func (h *handler) Get(c *gin.Context) {
	id := c.Param("id")
	people, err := h.useCase.GetOne(id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response := getPeopleResponse{
		Id:        id,
		Name:      people.Name,
		Surname:   people.Surname,
		Birthdate: people.Birthdate,
		CreatedAt: people.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
