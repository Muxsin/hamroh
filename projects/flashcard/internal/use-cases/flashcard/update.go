package flashcard

import "hamroh/flashcard/internal/models"

func (uc *useCase) Update(flashcard *models.Flashcard) error {
	return uc.service.Update(flashcard)
}
