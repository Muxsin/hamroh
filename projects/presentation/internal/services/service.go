package services

import "hamroh/presentation/internal/models"

type presentationRepository interface {
	Insert(presentation *models.Presentation) error
	GetAll() ([]*models.Presentation, error)
	GetById(id string) (*models.Presentation, error)
	Delete(id string) error
	Update(presentation *models.Presentation) error
}

type service struct {
	repository presentationRepository
}

func New(repository presentationRepository) *service {
	return &service{
		repository: repository,
	}
}
