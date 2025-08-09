package services

import "hamroh/recipe/internal/models"

type recipeRepository interface {
	Insert(recipe *models.Recipe) error
	GetAll() ([]*models.Recipe, error)
	GetById(id string) (*models.Recipe, error)
	Delete(id string) error
	Update(recipe *models.Recipe) error
}

type service struct {
	repository recipeRepository
}

func New(repository recipeRepository) *service {
	return &service{
		repository: repository,
	}
}
