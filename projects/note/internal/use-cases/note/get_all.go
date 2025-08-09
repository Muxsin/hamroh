package note

import "hamroh/note/internal/models"

func (uc *useCase) GetAll() ([]*models.Note, error) {
	note, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return note, nil
}
