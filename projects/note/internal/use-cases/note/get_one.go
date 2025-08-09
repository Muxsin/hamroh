package note

import "hamroh/note/internal/models"

func (uc *useCase) GetOne(id string) (*models.Note, error) {
	return uc.service.GetOne(id)
}
