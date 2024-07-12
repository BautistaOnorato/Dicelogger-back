package proficiencyXclass

import "github.com/proyecto-dnd/backend/internal/domain"

type ProficiencyXClassRepository interface {
	Create(proficiencyXClass domain.ProficiencyXClass) (domain.ProficiencyXClass, error)
	Delete(proficiencyXClass domain.ProficiencyXClass) error
}

type ProficiencyXClassService interface {
	Create(proficiencyXClass domain.ProficiencyXClass) (domain.ProficiencyXClass, error)
	Delete(proficiencyXClass domain.ProficiencyXClass) error
}
