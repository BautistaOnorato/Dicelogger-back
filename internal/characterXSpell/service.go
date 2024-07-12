package characterXspell

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	characterXSpellRepository RepositoryCharacterXSpell
}

func NewCharacterXSpellService(characterXSpellRepository RepositoryCharacterXSpell) ServiceCharacterXSpell {
	return &service{characterXSpellRepository: characterXSpellRepository}
}

// DeleteByCharacterDataId implements ServiceCharacterXSpell.
func (s *service) DeleteByCharacterDataId(id int) error {
	return s.characterXSpellRepository.DeleteByCharacterDataId(id)
}


func (s *service) Create(characterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error) {
	return s.characterXSpellRepository.Create(characterXSpell)
}

func (s *service) Delete(id int) error {
	return s.characterXSpellRepository.Delete(id)
}

func (s *service) DeleteParams(characterId int, spellId int) error {
	return s.characterXSpellRepository.DeleteParams(characterId, spellId)
}
