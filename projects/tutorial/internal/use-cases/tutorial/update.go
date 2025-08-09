package tutorial

import "hamroh/tutorial/internal/models"

func (uc *useCase) Update(tutorial *models.Tutorial) error {
	return uc.service.Update(tutorial)
}
