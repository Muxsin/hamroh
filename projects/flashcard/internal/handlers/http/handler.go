package http

import (
	"hamroh/flashcard/internal/models"

	"github.com/go-playground/validator/v10"
)

type flashcardUseCase interface {
	Create(flashcard *models.Flashcard) error
	GetAll() ([]*models.Flashcard, error)
	GetOne(id string) (*models.Flashcard, error)
	Delete(id string) error
	Update(flashcard *models.Flashcard) error
}

type handler struct {
	useCase flashcardUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase flashcardUseCase) *handler {
	return &handler{useCase: useCase}
}
