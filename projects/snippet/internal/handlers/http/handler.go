package http

import (
	"hamroh/snippet/internal/models"

	"github.com/go-playground/validator/v10"
)

type snippetUseCase interface {
	Create(snippet *models.Snippet) error
	GetAll() ([]*models.Snippet, error)
	GetOne(id string) (*models.Snippet, error)
	Delete(id string) error
	Update(snippet *models.Snippet) error
}

type handler struct {
	useCase snippetUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase snippetUseCase) *handler {
	return &handler{useCase: useCase}
}
