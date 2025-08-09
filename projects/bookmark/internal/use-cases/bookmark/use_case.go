package bookmark

import "hamroh/bookmark/internal/models"

type bookmarkService interface {
	Create(bookmark *models.Bookmark) error
	GetAll() ([]*models.Bookmark, error)
	GetOne(id string) (*models.Bookmark, error)
	Delete(id string) error
	Update(bookmark *models.Bookmark) error
}

type useCase struct {
	service bookmarkService
}

func New(service bookmarkService) *useCase {
	return &useCase{service: service}
}
