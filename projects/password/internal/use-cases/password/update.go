package password

import "kodnavis/password/internal/models"

func (uc *useCase) Update(password *models.Password) error {
	return uc.service.Update(password)
}
