package responses

import (
	"time"
	"todos-service/internal/models"
)

type TodoResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"Text"`
	Done      bool   `json:"body"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func NewTodoResponse(todo *models.Todo) *TodoResponse {
	return &TodoResponse{
		Id:        todo.ID,
		Text:      todo.Text,
		Done:      todo.Done,
		UserID:    todo.UserID,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
	}
}
