package people

import "kodnavis/people/internal/models"

func (uc *useCase) Update(people *models.People) error {
	return uc.service.Update(people)
}
