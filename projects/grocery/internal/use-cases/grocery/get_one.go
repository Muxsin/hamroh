package grocery

import "hamroh/grocery/internal/models"

func (uc *useCase) GetOne(id string) (*models.Grocery, error) {
	return uc.service.GetOne(id)
}
