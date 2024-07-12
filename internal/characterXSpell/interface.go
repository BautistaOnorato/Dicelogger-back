package characterXspell

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryCharacterXSpell interface {
	Create(characterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error)
	Delete(id int) error
	DeleteParams(characterId int, spellId int) error
	DeleteByCharacterDataId(id int) error
}

type ServiceCharacterXSpell interface {
	Create(characterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error)
	Delete(id int) error
	DeleteParams(characterId int, spellId int) error
	DeleteByCharacterDataId(id int) error
}
