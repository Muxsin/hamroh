package bookmark

import "hamroh/bookmark/internal/models"

func (uc *useCase) GetOne(id string) (*models.Bookmark, error) {
	return uc.service.GetOne(id)
}
