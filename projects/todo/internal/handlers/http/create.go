package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createTodoRequest struct {
	Title string `json:"title"`
}

type createTodoResponse struct {
	Id        uint      `json:"Id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (h *handler) Create(c *gin.Context) {
	var req createTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := h.useCase.Create(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := createTodoResponse{
		Id:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
