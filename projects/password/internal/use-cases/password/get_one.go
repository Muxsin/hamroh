package password

import "kodnavis/password/internal/models"

func (uc *useCase) GetOne(id string) (*models.Password, error) {
	return uc.service.GetOne(id)
}
