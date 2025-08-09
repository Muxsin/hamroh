package repositories

import "hamroh/note/internal/models"

func (r *repository) GetById(id string) (*models.Note, error) {
	var note models.Note
	if err := r.db.First(&note, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &note, nil
}
