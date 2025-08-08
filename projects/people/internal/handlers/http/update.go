package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type updatePeopleRequest struct {
	Name      string    `json:"name" validate:"required"`
	Surname   string    `json:"surname" validate:"required"`
	Birthdate time.Time `json:"birthdate" validate:"required"`
}

type updatePeopleResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Birthdate time.Time `json:"birthdate"`
	CreatedAt string    `json:"createdAt"`
}

func (h *handler) Update(c *gin.Context) {
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

	var req updatePeopleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	people.Name = req.Name
	people.Surname = req.Surname
	people.Birthdate = req.Birthdate

	if err := h.useCase.Update(people); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	response := updatePeopleResponse{
		Id:        id,
		Name:      people.Name,
		Surname:   people.Surname,
		Birthdate: people.Birthdate,
		CreatedAt: people.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
