package characterXproficiency

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repository CharacterXProficiencyRepository
}

func NewCharacterXProficiencyService(repository CharacterXProficiencyRepository) CharacterXProficiencyService {
	return &service{repository: repository}
}

// DeleteByCharacterDataId implements CharacterXProficiencyService.
func (s *service) DeleteByCharacterDataId(id int) error {
	return s.repository.DeleteByCharacterDataId(id)
}


func (s *service) Create(characterXProficiency domain.CharacterXProficiency) (domain.CharacterXProficiency, error) {
	createdCharacterXProficiency, err := s.repository.Create(characterXProficiency)
	if err != nil {
		return domain.CharacterXProficiency{}, err
	}

	return createdCharacterXProficiency, nil
}

func (s *service) Delete(characterXProficiencyId int) error {
	return s.repository.Delete(characterXProficiencyId)
}
