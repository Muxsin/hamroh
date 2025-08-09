package http

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getPasswordResponse struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) Get(c *gin.Context) {
	id := c.Param("id")
	password, err := h.useCase.GetOne(id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response := getPasswordResponse{
		Id:        id,
		Text:      password.Text,
		CreatedAt: password.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
