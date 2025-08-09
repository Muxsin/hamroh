package recipe

import "hamroh/recipe/internal/models"

type recipeService interface {
	Create(recipe *models.Recipe) error
	GetAll() ([]*models.Recipe, error)
	GetOne(id string) (*models.Recipe, error)
	Delete(id string) error
	Update(recipe *models.Recipe) error
}

type useCase struct {
	service recipeService
}

func New(service recipeService) *useCase {
	return &useCase{service: service}
}
