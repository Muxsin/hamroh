package people

import "kodnavis/people/internal/models"

type peopleService interface {
	Create(people *models.People) error
}

type useCase struct {
	service peopleService
}

func New(service peopleService) *useCase {
	return &useCase{service: service}
}
