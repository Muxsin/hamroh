package services

import "hamroh/snippet/internal/models"

type snippetRepository interface {
	Insert(snippet *models.Snippet) error
	GetAll() ([]*models.Snippet, error)
	GetById(id string) (*models.Snippet, error)
	Delete(id string) error
	Update(snippet *models.Snippet) error
}

type service struct {
	repository snippetRepository
}

func New(repository snippetRepository) *service {
	return &service{
		repository: repository,
	}
}
