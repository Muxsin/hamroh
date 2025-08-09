package recipe

import "hamroh/recipe/internal/models"

func (uc *useCase) Update(recipe *models.Recipe) error {
	return uc.service.Update(recipe)
}
