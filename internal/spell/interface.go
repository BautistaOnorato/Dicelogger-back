package spell

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RepositorySpell interface {
	Create(spell dto.SpellDto) (domain.Spell, error)
	GetAll() ([]domain.Spell, error)
	GetById(id int) (domain.Spell, error)
	GetByCharacterDataId(id int) ([]domain.Spell, error)
	GetByClassId(classId int) ([]domain.Spell, error)
	Update(spell dto.SpellDto, id int) (domain.Spell, error)
	Delete(id int) error
}

type ServiceSpell interface {
	Create(spell dto.SpellDto) (domain.Spell, error)
	GetAll() ([]domain.Spell, error)
	GetById(id int) (domain.Spell, error)
	GetByCharacterDataId(id int) ([]domain.Spell, error)
	GetByClassId(classId int) ([]domain.Spell, error)
	Update(spell dto.SpellDto, id int) (domain.Spell, error)
	Delete(id int) error
}
