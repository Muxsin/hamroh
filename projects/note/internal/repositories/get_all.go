package repositories

import "hamroh/note/internal/models"

func (r *repository) GetAll() ([]*models.Note, error) {
	var notes []*models.Note

	if err := r.db.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}
