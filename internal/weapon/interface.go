package weapon

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryWeapon interface {
	Create(weapon domain.Weapon) (domain.Weapon, error)
	GetAll() ([]domain.Weapon, error)
	GetByCampaignId(campaignId int) ([]domain.Weapon, error)
	GetAllGeneric() ([]domain.Weapon, error)
	GetById(id int) (domain.Weapon, error)
	Update(weapon domain.Weapon) (domain.Weapon, error)
	Delete(id int) error
}

type ServiceWeapon interface {
	Create(weapon domain.Weapon) (domain.Weapon, error)
	GetAll() ([]domain.Weapon, error)
	GetByCampaignId(campaignId int) ([]domain.Weapon, error)
	GetAllGeneric()	([]domain.Weapon, error)
	GetById(id int) (domain.Weapon, error)
	Update(item domain.Weapon) (domain.Weapon, error)
	Delete(id int) error
}