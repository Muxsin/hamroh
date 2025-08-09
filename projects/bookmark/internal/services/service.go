package services

import "hamroh/bookmark/internal/models"

type bookmarkRepository interface {
	Insert(bookmark *models.Bookmark) error
	GetAll() ([]*models.Bookmark, error)
	GetById(id string) (*models.Bookmark, error)
	Delete(id string) error
	Update(bookmark *models.Bookmark) error
}

type service struct {
	repository bookmarkRepository
}

func New(repository bookmarkRepository) *service {
	return &service{
		repository: repository,
	}
}
