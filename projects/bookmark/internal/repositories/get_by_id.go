package repositories

import "hamroh/bookmark/internal/models"

func (r *repository) GetById(id string) (*models.Bookmark, error) {
	var bookmark models.Bookmark
	if err := r.db.First(&bookmark, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &bookmark, nil
}
