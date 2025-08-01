package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ListTodoResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) List(c *gin.Context) {
	todos, err := h.UseCase.List()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var response []ListTodoResponse
	for _, todo := range todos {
		response = append(response, ListTodoResponse{
			Id:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
