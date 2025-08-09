package tutorial

import "hamroh/tutorial/internal/models"

func (uc *useCase) GetAll() ([]*models.Tutorial, error) {
	tutorial, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return tutorial, nil
}
