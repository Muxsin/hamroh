package http

import (
	"hamroh/recipe/internal/models"

	"github.com/go-playground/validator/v10"
)

type recipeUseCase interface {
	Create(recipe *models.Recipe) error
	GetAll() ([]*models.Recipe, error)
	GetOne(id string) (*models.Recipe, error)
	Delete(id string) error
	Update(recipe *models.Recipe) error
}

type handler struct {
	useCase recipeUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase recipeUseCase) *handler {
	return &handler{useCase: useCase}
}
