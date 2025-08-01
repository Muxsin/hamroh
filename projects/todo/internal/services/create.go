package services

import (
	"errors"
	"hamroh-todo/internal/models"
)

func (s *service) Create(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	return s.repository.Insert(todo)
}
