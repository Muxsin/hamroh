package bookmark

import "hamroh/bookmark/internal/models"

func (uc *useCase) Create(bookmark *models.Bookmark) error {
	return uc.service.Create(bookmark)
}
