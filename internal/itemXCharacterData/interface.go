package itemxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryItemXCharacterData interface {
	Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	GetAll() ([]domain.ItemXCharacterData, error)
	GetById(id int) (domain.ItemXCharacterData, error)
	GetByCharacterDataId(id int) ([]domain.ItemXCharacterData, error)
	Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	UpdateOwnership(itemXCharacterData domain.ItemXCharacterData) (error)
	Delete(id int) error
	DeleteByCharacterDataId(id int) error
}

type ServiceItemXCharacterData interface {
	Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	GetAll() ([]domain.ItemXCharacterData, error)
	GetById(id int) (domain.ItemXCharacterData, error)
	GetByCharacterDataId(id int) ([]domain.ItemXCharacterData, error)
	Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	UpdateOwnership(itemXCharacterData domain.ItemXCharacterData) (error)
	Delete(id int) error
	DeleteByCharacterDataId(id int) error
}