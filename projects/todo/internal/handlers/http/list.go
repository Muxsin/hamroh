package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type listTodoResponse struct {
	id        uint
	title     string
	completed bool
	createdAt time.Time `json:"created_at"`
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
			id:        todo.ID,
			title:     todo.Title,
			completed: todo.Completed,
			createdAt: todo.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
