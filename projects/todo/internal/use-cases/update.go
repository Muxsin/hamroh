package use_cases

import "hamroh-todo/internal/models"

func (uc *useCase) Update(todo *models.Todo) error {
	return uc.service.Update(todo)
}
