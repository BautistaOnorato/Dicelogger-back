package proficiency

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RepositoryProficiency interface {
	Create(proficiencyDto dto.ProficiencyDto) (domain.Proficiency, error)
	GetAll() ([]domain.Proficiency, error)
	GetById(id int) (domain.Proficiency, error)
	GetByCharacterDataId(id int) ([]domain.Proficiency, error)
	Update(proficiencyDto dto.ProficiencyDto, id int) (domain.Proficiency, error)
	Delete(id int) error
}

type ProficiencyService interface {
	Create(proficiencyDto dto.ProficiencyDto) (domain.Proficiency, error)
	GetAll() ([]domain.Proficiency, error)
	GetById(id int) (domain.Proficiency, error)
	GetByCharacterDataId(id int) ([]domain.Proficiency, error)
	Update(proficiencyDto dto.ProficiencyDto, id int) (domain.Proficiency, error)
	Delete(id int) error
}