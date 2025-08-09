package flashcard

import "hamroh/flashcard/internal/models"

type flashcardService interface {
	Create(flashcard *models.Flashcard) error
	GetAll() ([]*models.Flashcard, error)
	GetOne(id string) (*models.Flashcard, error)
	Delete(id string) error
	Update(flashcard *models.Flashcard) error
}

type useCase struct {
	service flashcardService
}

func New(service flashcardService) *useCase {
	return &useCase{service: service}
}
