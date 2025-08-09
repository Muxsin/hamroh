package grocery

import "hamroh/grocery/internal/models"

func (uc *useCase) Create(grocery *models.Grocery) error {
	return uc.service.Create(grocery)
}
