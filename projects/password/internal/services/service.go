package services

import "kodnavis/password/internal/models"

type passwordRepository interface {
	Insert(password *models.Password) error
	GetAll() ([]*models.Password, error)
	GetById(id string) (*models.Password, error)
	Delete(id string) error
	Update(password *models.Password) error
}

type service struct {
	repository passwordRepository
}

func New(repository passwordRepository) *service {
	return &service{
		repository: repository,
	}
}
