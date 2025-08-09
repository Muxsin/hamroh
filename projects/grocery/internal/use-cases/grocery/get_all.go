package grocery

import "hamroh/grocery/internal/models"

func (uc *useCase) GetAll() ([]*models.Grocery, error) {
	grocery, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return grocery, nil
}
