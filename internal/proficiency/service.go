package proficiency

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repository RepositoryProficiency
}

func NewProficiencyService(repository RepositoryProficiency) ProficiencyService {
	return &service{repository: repository}
}

func (s *service) Create(proficiencyDto dto.ProficiencyDto) (domain.Proficiency, error) {
	createdProficiency, err := s.repository.Create(proficiencyDto)
	if err != nil {
		return domain.Proficiency{}, err
	}

	return createdProficiency, nil
}

func (s *service) GetAll() ([]domain.Proficiency, error) {
	serviceList, err := s.repository.GetAll()
	if err != nil {
		return []domain.Proficiency{}, err
	}

	return serviceList, nil
}

func (s *service) GetById(id int) (domain.Proficiency, error) {
	proficiency, err := s.repository.GetById(id)
	if err != nil {
		return domain.Proficiency{}, err
	}

	return proficiency, nil
}

func (s *service) Update(proficiencyDto dto.ProficiencyDto, id int) (domain.Proficiency, error) {
	updatedProficiency, err := s.repository.Update(proficiencyDto, id)
	if err != nil {
		return domain.Proficiency{}, err
	}

	return updatedProficiency, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) GetByCharacterDataId(characterId int) ([]domain.Proficiency, error) {
	proficiencyList, err := s.repository.GetByCharacterDataId(characterId)
	if err != nil {
		return []domain.Proficiency{}, err
	}
	return proficiencyList, nil
}
