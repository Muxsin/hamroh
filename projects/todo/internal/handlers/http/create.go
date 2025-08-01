package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type CreateTodoResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := h.UseCase.Create(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := CreateTodoResponse{
		Id:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
