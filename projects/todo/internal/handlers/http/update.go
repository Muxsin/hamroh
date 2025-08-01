package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateTodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *handler) Update(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.useCase.GetById(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	var request updateTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.Title = request.Title
	todo.Completed = request.Completed

	if err := h.useCase.Update(&todo); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
