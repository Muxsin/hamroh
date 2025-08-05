package services

func (s *service) Delete(id string) error {
	return s.repository.Delete(id)
}
