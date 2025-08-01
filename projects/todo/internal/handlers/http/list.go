package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type listTodoResponse struct {
	Id        uint      `json:"Id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	todos, err := h.useCase.List()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var response []listTodoResponse
	for _, todo := range todos {
		response = append(response, listTodoResponse{
			Id:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
