package todo

import "hamroh-todo/internal/models"

func (uc *useCase) GetAll() ([]models.Todo, error) {
	todos, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
