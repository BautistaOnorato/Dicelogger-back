package class

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RepositoryCharacterClass interface {
	Create(classDto dto.ClassDto) (domain.Class, error)
	GetAll()([]domain.Class, error)
	GetById(id int)(domain.Class, error)
	Update(classDto dto.ClassDto,id int)(domain.Class, error)
	Delete(id int) error
}

type ClassService interface {
	Create(classDto dto.ClassDto) (domain.Class, error)
	GetAll()([]domain.Class, error)
	GetById(id int)(domain.Class, error)
	Update(classDto dto.ClassDto,id int)(domain.Class, error)
	Delete(id int) error
}