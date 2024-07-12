package classXspell

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryClassXSpell interface {
	Create(classXClassData domain.ClassXSpell) (domain.ClassXSpell, error)
	Delete(classXClassData domain.ClassXSpell) error
}

type ServiceClassXSpell interface {
	Create(classXClassData domain.ClassXSpell) (domain.ClassXSpell, error)
	Delete(classXClassData domain.ClassXSpell) error
}