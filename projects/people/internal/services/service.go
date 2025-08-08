package services

import "kodnavis/people/internal/models"

type peopleRepository interface {
	Insert(people *models.People) error
	GetAll() ([]*models.People, error)
	GetById(id string) (*models.People, error)
}

type service struct {
	repository peopleRepository
}

func New(repository peopleRepository) *service {
	return &service{
		repository: repository,
	}
}
