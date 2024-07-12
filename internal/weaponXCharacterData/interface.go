package weaponxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryWeaponXCharacterData interface {
	Create(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	GetAll() ([]domain.WeaponXCharacterData, error)
	GetById(id int) (domain.WeaponXCharacterData, error)
	GetByCharacterDataId(id int) ([]domain.WeaponXCharacterData, error)
	Update(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	UpdateOwnership(WeaponXCharacterData domain.WeaponXCharacterData) (error)
	Delete(id int) error
	DeleteByCharacterDataId(id int) error
}

type ServiceWeaponXCharacterData interface {
	Create(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	GetAll() ([]domain.WeaponXCharacterData, error)
	GetById(id int) (domain.WeaponXCharacterData, error)
	GetByCharacterDataId(id int) ([]domain.WeaponXCharacterData, error)
	Update(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	UpdateOwnership(WeaponXCharacterData domain.WeaponXCharacterData) (error)
	Delete(id int) error
	DeleteByCharacterDataId(id int) error
}