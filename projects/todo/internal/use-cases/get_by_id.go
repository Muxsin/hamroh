package use_cases

import "hamroh-todo/internal/models"

func (uc *useCase) GetById(id string) (models.Todo, error) {
	todo, err := uc.service.GetById(id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}
