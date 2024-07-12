package skill

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositorySkill interface {
	Create(skill domain.Skill) (domain.Skill, error)
	GetAll() ([]domain.Skill, error)
	GetById(id int) (domain.Skill, error)
	GetByCharacterId(characterId int) ([]domain.Skill, error)
	GetByClassId(classId int) ([]domain.Skill, error)
	Update(skill domain.Skill) (domain.Skill, error)
	Delete(id int) error
}

type ServiceSkill interface {
	Create(skill domain.Skill) (domain.Skill, error)
	GetAll() ([]domain.Skill, error)
	GetById(id int) (domain.Skill, error)
	GetByCharacterId(characterId int) ([]domain.Skill, error)
	GetByClassId(classId int) ([]domain.Skill, error)
	Update(skill domain.Skill) (domain.Skill, error)
	Delete(id int) error
}
