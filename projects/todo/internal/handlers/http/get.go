package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type getTodoResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *handler) Get(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.useCase.GetOne(id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response := getTodoResponse{
		Id:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}
