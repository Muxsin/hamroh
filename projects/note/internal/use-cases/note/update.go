package note

import "hamroh/note/internal/models"

func (uc *useCase) Update(note *models.Note) error {
	return uc.service.Update(note)
}
