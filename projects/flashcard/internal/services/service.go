package services

import "hamroh/flashcard/internal/models"

type flashcardRepository interface {
	Insert(flashcard *models.Flashcard) error
	GetAll() ([]*models.Flashcard, error)
	GetById(id string) (*models.Flashcard, error)
	Delete(id string) error
	Update(flashcard *models.Flashcard) error
}

type service struct {
	repository flashcardRepository
}

func New(repository flashcardRepository) *service {
	return &service{
		repository: repository,
	}
}
