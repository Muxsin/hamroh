package repositories

import "hamroh/flashcard/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Flashcard{}, "id = ?", id).Error
}
