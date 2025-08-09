package snippet

import "hamroh/snippet/internal/models"

func (uc *useCase) Create(snippet *models.Snippet) error {
	return uc.service.Create(snippet)
}
