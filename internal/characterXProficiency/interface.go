package characterXproficiency

import "github.com/proyecto-dnd/backend/internal/domain"

type CharacterXProficiencyRepository interface {
	Create(characterXProficiency domain.CharacterXProficiency) (domain.CharacterXProficiency, error)
	Delete(characterXProficiencyId int) error
	DeleteByCharacterDataId(id int) error
}

type CharacterXProficiencyService interface {
	Create(characterXProficiency domain.CharacterXProficiency) (domain.CharacterXProficiency, error)
	Delete(characterXProficiencyId int) error
	DeleteByCharacterDataId(id int) error
}
