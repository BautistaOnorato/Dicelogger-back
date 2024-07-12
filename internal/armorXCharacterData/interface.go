package armorXCharacterData

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryArmorXCharacterData interface {
	CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
	DeleteArmorXCharacterData(id int) error
	DeleteByCharacterDataIdArmor(id int) error
	GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error)
	GetByIdArmorXCharacterData(id int) (domain.ArmorXCharacterData, error)
	GetByCharacterDataIdArmor(id int) ([]domain.ArmorXCharacterData, error)
	UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
	UpdateOwnership(armorXCharacterData domain.ArmorXCharacterData) (error)
}

type ServiceArmorXCharacterData interface {
	CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
	DeleteArmorXCharacterData(id int) error
	DeleteByCharacterDataIdArmor(id int) error
	GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error)
	GetByIdArmorXCharacterData(id int) (domain.ArmorXCharacterData, error)
	GetByCharacterDataIdArmor(id int) ([]domain.ArmorXCharacterData, error)
	UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
	UpdateOwnership(armorXCharacterData domain.ArmorXCharacterData) (error)
}
