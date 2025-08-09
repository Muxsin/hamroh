package flashcard

import "hamroh/flashcard/internal/models"

func (uc *useCase) Create(flashcard *models.Flashcard) error {
	return uc.service.Create(flashcard)
}
