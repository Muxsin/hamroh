package repositories

import "hamroh/bookmark/internal/models"

func (r *repository) GetAll() ([]*models.Bookmark, error) {
	var bookmarks []*models.Bookmark

	if err := r.db.Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}
