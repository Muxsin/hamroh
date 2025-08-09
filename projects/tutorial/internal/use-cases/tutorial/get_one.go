package tutorial

import "hamroh/tutorial/internal/models"

func (uc *useCase) GetOne(id string) (*models.Tutorial, error) {
	return uc.service.GetOne(id)
}
