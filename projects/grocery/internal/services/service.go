package services

import "hamroh/grocery/internal/models"

type groceryRepository interface {
	Insert(grocery *models.Grocery) error
	GetAll() ([]*models.Grocery, error)
	GetById(id string) (*models.Grocery, error)
	Delete(id string) error
	Update(grocery *models.Grocery) error
}

type service struct {
	repository groceryRepository
}

func New(repository groceryRepository) *service {
	return &service{
		repository: repository,
	}
}
