package grocery

import "hamroh/grocery/internal/models"

func (uc *useCase) Update(grocery *models.Grocery) error {
	return uc.service.Update(grocery)
}
