package handlers

import (
	"github.com/gin-gonic/gin"
	"kodnavis/people/internal/models"
	"log"
	"net/http"
	"time"
)

type createPeopleRequest struct {
	Name     string    `json:"name" validate:"required"`
	Surname  string    `json:"surname" validate:"required"`
	Birthday time.Time `json:"birthday" validate:"required"`
}

func (h *handler) Create(c *gin.Context) {
	var req createPeopleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	people := &models.People{
		Name:      req.Name,
		Surname:   req.Surname,
		Birthdate: req.Birthday,
	}

	if err := h.useCase.Create(people); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
