package recipe

import "hamroh/recipe/internal/models"

func (uc *useCase) GetAll() ([]*models.Recipe, error) {
	recipe, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
