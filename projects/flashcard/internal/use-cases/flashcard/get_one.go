package flashcard

import "hamroh/flashcard/internal/models"

func (uc *useCase) GetOne(id string) (*models.Flashcard, error) {
	return uc.service.GetOne(id)
}
