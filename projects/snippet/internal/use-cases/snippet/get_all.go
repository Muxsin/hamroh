package snippet

import "hamroh/snippet/internal/models"

func (uc *useCase) GetAll() ([]*models.Snippet, error) {
	snippet, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return snippet, nil
}
