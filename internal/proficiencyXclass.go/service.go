package proficiencyXclass

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repository ProficiencyXClassRepository
}

func NewProficiencyXClassService(repository ProficiencyXClassRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(proficiencyXClass domain.ProficiencyXClass) (domain.ProficiencyXClass, error) {
	createdProficiencyXClass, err := s.repository.Create(proficiencyXClass)
	if err != nil {
		return domain.ProficiencyXClass{}, err
	}

	return createdProficiencyXClass, nil
}

func (s *service) Delete(proficiencyXClass domain.ProficiencyXClass) error {
	return s.repository.Delete(proficiencyXClass)
}
