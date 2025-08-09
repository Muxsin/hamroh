package tutorial

import "hamroh/tutorial/internal/models"

func (uc *useCase) Create(tutorial *models.Tutorial) error {
	return uc.service.Create(tutorial)
}
