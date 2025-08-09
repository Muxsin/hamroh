package tutorial

import "hamroh/tutorial/internal/models"

type tutorialService interface {
	Create(tutorial *models.Tutorial) error
	GetAll() ([]*models.Tutorial, error)
	GetOne(id string) (*models.Tutorial, error)
	Delete(id string) error
	Update(tutorial *models.Tutorial) error
}

type useCase struct {
	service tutorialService
}

func New(service tutorialService) *useCase {
	return &useCase{service: service}
}
