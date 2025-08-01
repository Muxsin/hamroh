package repositories

import "hamroh-todo/internal/models"

func (r *repository) Update(todo *models.Todo) error {
	return r.db.Updates(todo).Error
}
