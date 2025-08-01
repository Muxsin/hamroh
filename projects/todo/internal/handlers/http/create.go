package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createTodoRequest struct {
	title string
}

type createTodoResponse struct {
	id        uint
	title     string
	completed bool
	createdAt time.Time `json:"created_at"`
}

func (h *handler) Create(c *gin.Context) {
	var req createTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := h.useCase.Create(req.title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := createTodoResponse{
		id:        todo.ID,
		title:     todo.Title,
		completed: todo.Completed,
		createdAt: todo.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
