package recipe

import "hamroh/recipe/internal/models"

func (uc *useCase) GetOne(id string) (*models.Recipe, error) {
	return uc.service.GetOne(id)
}
