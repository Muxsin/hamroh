package use_cases

import "hamroh-todo/internal/models"

func (uc *useCase) List() ([]models.Todo, error) {
	todos, err := uc.service.List()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
