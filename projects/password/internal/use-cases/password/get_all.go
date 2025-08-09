package password

import "kodnavis/password/internal/models"

func (uc *useCase) GetAll() ([]*models.Password, error) {
	password, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return password, nil
}
