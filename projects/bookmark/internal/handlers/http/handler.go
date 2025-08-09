package http

import (
	"hamroh/bookmark/internal/models"

	"github.com/go-playground/validator/v10"
)

type bookmarkUseCase interface {
	Create(bookmark *models.Bookmark) error
	GetAll() ([]*models.Bookmark, error)
	GetOne(id string) (*models.Bookmark, error)
	Delete(id string) error
	Update(bookmark *models.Bookmark) error
}

type handler struct {
	useCase bookmarkUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase bookmarkUseCase) *handler {
	return &handler{useCase: useCase}
}
