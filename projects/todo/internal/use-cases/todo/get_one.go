package todo

import "hamroh-todo/internal/models"

func (uc *useCase) GetOne(id string) (models.Todo, error) {
	todo, err := uc.service.GetOne(id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}
