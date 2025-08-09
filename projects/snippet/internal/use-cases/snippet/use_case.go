package snippet

import "hamroh/snippet/internal/models"

type snippetService interface {
	Create(snippet *models.Snippet) error
	GetAll() ([]*models.Snippet, error)
	GetOne(id string) (*models.Snippet, error)
	Delete(id string) error
	Update(snippet *models.Snippet) error
}

type useCase struct {
	service snippetService
}

func New(service snippetService) *useCase {
	return &useCase{service: service}
}
