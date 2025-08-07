package people

import "kodnavis/people/internal/models"

func (uc *useCase) GetAll() ([]*models.People, error) {
	people, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return people, nil
}
