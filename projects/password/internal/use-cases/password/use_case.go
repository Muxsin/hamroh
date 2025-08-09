package password

import "kodnavis/password/internal/models"

type passwordService interface {
	Create(password *models.Password) error
	GetAll() ([]*models.Password, error)
	GetOne(id string) (*models.Password, error)
	Delete(id string) error
	Update(password *models.Password) error
}

type useCase struct {
	service passwordService
}

func New(service passwordService) *useCase {
	return &useCase{service: service}
}
