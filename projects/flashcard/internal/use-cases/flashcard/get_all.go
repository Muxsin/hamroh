package flashcard

import "hamroh/flashcard/internal/models"

func (uc *useCase) GetAll() ([]*models.Flashcard, error) {
	flashcard, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return flashcard, nil
}
