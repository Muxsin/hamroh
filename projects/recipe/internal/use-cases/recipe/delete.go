package recipe

func (uc *useCase) Delete(id string) error {
	return uc.service.Delete(id)
}
