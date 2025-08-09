package services

func (r *service) Delete(id string) error {
	return r.repository.Delete(id)
}
