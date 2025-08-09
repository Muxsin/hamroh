package bookmark

import "hamroh/bookmark/internal/models"

func (uc *useCase) GetAll() ([]*models.Bookmark, error) {
	bookmark, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}
