package repositories

import "hamroh-todo/internal/models"

func (r *repository) Insert(todo *models.Todo) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(todo).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
