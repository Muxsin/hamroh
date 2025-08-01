package http

import (
	"github.com/gin-gonic/gin"
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
	todo, err := h.useCase.GetById(id)
	if err != nil {
		c.Status(http.StatusNotFound)
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
