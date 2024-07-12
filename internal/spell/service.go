package spell

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repository RepositorySpell
}

func NewSpellService(repository RepositorySpell) ServiceSpell {
	return &service{repository: repository}
}

// (name, description, range, ritual, duration, concentration, casting_time, level, damage_type, difficulty_class, aoe, school)
func (s *service) Create(spellDto dto.SpellDto) (domain.Spell, error) {
	createdSpell, err := s.repository.Create(spellDto)
	if err != nil {
		return domain.Spell{}, err
	}

	return createdSpell, nil
}

func (s *service) GetAll() ([]domain.Spell, error) {
	spells, err := s.repository.GetAll()
	if err != nil {
		return []domain.Spell{}, err
	}

	return spells, nil
}

func (s *service) GetById(id int) (domain.Spell, error) {
	spell, err := s.repository.GetById(id)
	if err != nil {
		return domain.Spell{}, err
	}

	return spell, nil
}

func (s *service) Update(spellDto dto.SpellDto, id int) (domain.Spell, error) {
	return s.repository.Update(spellDto, id)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) GetByCharacterDataId(characterId int) ([]domain.Spell, error) {
	spellList, err := s.repository.GetByCharacterDataId(characterId)
	if err != nil {
		return []domain.Spell{}, err
	}
	return spellList, nil
}

func (s *service) GetByClassId(classId int) ([]domain.Spell, error) {
	return s.repository.GetByClassId(classId)
}
