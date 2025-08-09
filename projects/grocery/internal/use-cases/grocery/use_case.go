package grocery

import "hamroh/grocery/internal/models"

type groceryService interface {
	Create(grocery *models.Grocery) error
	GetAll() ([]*models.Grocery, error)
	GetOne(id string) (*models.Grocery, error)
	Delete(id string) error
	Update(grocery *models.Grocery) error
}

type useCase struct {
	service groceryService
}

func New(service groceryService) *useCase {
	return &useCase{service: service}
}
