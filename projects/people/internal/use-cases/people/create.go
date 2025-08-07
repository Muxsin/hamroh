package people

import "kodnavis/people/internal/models"

func (uc *useCase) Create(people *models.People) error {
	return uc.service.Create(people)
}
