package bookmark

import "hamroh/bookmark/internal/models"

func (uc *useCase) Update(bookmark *models.Bookmark) error {
	return uc.service.Update(bookmark)
}
