package services

import (
	"errors"
	"fmt"
	"hamroh-todo/internal/models"
)

func (s *service) Create(todo *models.Todo) error {
	fmt.Println(todo.Title)
	if todo.Title == "" {
		return errors.New("title is required")
	}

	return s.repository.Insert(todo)
}
