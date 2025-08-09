package services

import "hamroh/note/internal/models"

type noteRepository interface {
	Insert(note *models.Note) error
	GetAll() ([]*models.Note, error)
	GetById(id string) (*models.Note, error)
	Delete(id string) error
	Update(note *models.Note) error
}

type service struct {
	repository noteRepository
}

func New(repository noteRepository) *service {
	return &service{
		repository: repository,
	}
}
