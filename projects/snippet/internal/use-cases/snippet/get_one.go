package snippet

import "hamroh/snippet/internal/models"

func (uc *useCase) GetOne(id string) (*models.Snippet, error) {
	return uc.service.GetOne(id)
}
