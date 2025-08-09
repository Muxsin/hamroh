package presentation

import "hamroh/presentation/internal/models"

type presentationService interface {
	Create(presentation *models.Presentation) error
	GetAll() ([]*models.Presentation, error)
	GetOne(id string) (*models.Presentation, error)
	Delete(id string) error
	Update(presentation *models.Presentation) error
}

type useCase struct {
	service presentationService
}

func New(service presentationService) *useCase {
	return &useCase{service: service}
}
