package note

import "hamroh/note/internal/models"

type noteService interface {
	Create(note *models.Note) error
	GetAll() ([]*models.Note, error)
	GetOne(id string) (*models.Note, error)
	Delete(id string) error
	Update(note *models.Note) error
}

type useCase struct {
	service noteService
}

func New(service noteService) *useCase {
	return &useCase{service: service}
}
