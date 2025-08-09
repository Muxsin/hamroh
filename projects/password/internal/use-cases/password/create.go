package password

import "kodnavis/password/internal/models"

func (uc *useCase) Create(password *models.Password) error {
	return uc.service.Create(password)
}
