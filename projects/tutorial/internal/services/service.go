package services

import "hamroh/tutorial/internal/models"

type tutorialRepository interface {
	Insert(tutorial *models.Tutorial) error
	GetAll() ([]*models.Tutorial, error)
	GetById(id string) (*models.Tutorial, error)
	Delete(id string) error
	Update(tutorial *models.Tutorial) error
}

type service struct {
	repository tutorialRepository
}

func New(repository tutorialRepository) *service {
	return &service{
		repository: repository,
	}
}
