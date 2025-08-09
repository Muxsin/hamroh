package repositories

import "hamroh/note/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Note{}, "id = ?", id).Error
}
