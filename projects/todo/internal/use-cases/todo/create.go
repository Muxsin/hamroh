package todo

import "hamroh-todo/internal/models"

func (uc *useCase) Create(title string) (*models.Todo, error) {
	todo := &models.Todo{
		Title:     title,
		Completed: false,
	}

	err := uc.service.Create(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
