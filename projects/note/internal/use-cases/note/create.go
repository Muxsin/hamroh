package note

import "hamroh/note/internal/models"

func (uc *useCase) Create(note *models.Note) error {
	return uc.service.Create(note)
}
