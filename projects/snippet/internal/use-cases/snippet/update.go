package snippet

import "hamroh/snippet/internal/models"

func (uc *useCase) Update(snippet *models.Snippet) error {
	return uc.service.Update(snippet)
}
