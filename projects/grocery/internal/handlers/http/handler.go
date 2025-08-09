package http

import (
	"hamroh/grocery/internal/models"

	"github.com/go-playground/validator/v10"
)

type groceryUseCase interface {
	Create(grocery *models.Grocery) error
	GetAll() ([]*models.Grocery, error)
	GetOne(id string) (*models.Grocery, error)
	Delete(id string) error
	Update(grocery *models.Grocery) error
}

type handler struct {
	useCase groceryUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase groceryUseCase) *handler {
	return &handler{useCase: useCase}
}
