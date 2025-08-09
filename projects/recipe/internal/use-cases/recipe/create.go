package recipe

import "hamroh/recipe/internal/models"

func (uc *useCase) Create(recipe *models.Recipe) error {
	return uc.service.Create(recipe)
}
