package people

import "kodnavis/people/internal/models"

func (uc *useCase) GetOne(id string) (*models.People, error) {
	return uc.service.GetOne(id)
}
