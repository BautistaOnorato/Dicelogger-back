package characterxattackevent

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterXAttackEventRepository interface {
	GetAll() ([]domain.CharacterXAttackEvent, error)
	GetById(id int) (domain.CharacterXAttackEvent, error)
	GetByCharacterId(characterId int) ([]domain.CharacterXAttackEvent, error)
	GetByEventId(attackEventId int) ([]domain.CharacterXAttackEvent, error)
	Create(CharacterXAttackEvent domain.CharacterXAttackEvent) (domain.CharacterXAttackEvent, error)
	Delete(id int) error
}

type CharacterXAttackEventService interface {
	GetAll() ([]domain.CharacterXAttackEvent, error)
	GetById(id int) (domain.CharacterXAttackEvent, error)
	GetByCharacterId(characterId int) ([]domain.CharacterXAttackEvent, error)
	GetByEventId(attackEventId int) ([]domain.CharacterXAttackEvent, error)
	Create(CharacterXAttackEvent dto.CharacterXAttackEventDto) (domain.CharacterXAttackEvent, error)
	Delete(id int) error
}
