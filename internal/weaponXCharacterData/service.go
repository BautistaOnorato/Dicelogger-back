package weaponxcharacterdata

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/weapon"
)

type service struct {
	weaponXCharacterDataRepo RepositoryWeaponXCharacterData
    weaponService weapon.ServiceWeapon
}

// Create implements ServiceItemXTableCharacter.
func (s *service) Create(weaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error) {
	newWeaponRelationship, err := s.weaponXCharacterDataRepo.Create(weaponXCharacterData)
	if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
    newWeaponRelationship.Weapon, err = s.weaponService.GetById(weaponXCharacterData.Weapon.Weapon_Id)
    if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
	return newWeaponRelationship, nil
}

// Delete implements ServiceItemXTableCharacter.
func (s *service) Delete(id int) error {
	err := s.weaponXCharacterDataRepo.Delete(id)
    if err!= nil {
        return err
    }
    return nil
}

// DeleteByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) DeleteByCharacterDataId(id int) error {
	err := s.weaponXCharacterDataRepo.DeleteByCharacterDataId(id)
    if err!= nil {
        return err
    }
    return nil
}

// GetAll implements ServiceItemXTableCharacter.
func (s *service) GetAll() ([]domain.WeaponXCharacterData, error) {
	weaponRelationships, err := s.weaponXCharacterDataRepo.GetAll()
	if err != nil {
		return []domain.WeaponXCharacterData{}, err
	}
	return weaponRelationships, nil
}

// GetByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) GetByCharacterDataId(id int) ([]domain.WeaponXCharacterData, error) {
	weaponRelationships, err := s.weaponXCharacterDataRepo.GetByCharacterDataId(id)
    if err!= nil {
        return []domain.WeaponXCharacterData{}, err
    }
    return weaponRelationships, nil
}

// GetById implements ServiceItemXTableCharacter.
func (s *service) GetById(id int) (domain.WeaponXCharacterData, error) {
	weaponRelationship, err := s.weaponXCharacterDataRepo.GetById(id)
    if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
    return weaponRelationship, nil
}

// Update implements ServiceItemXTableCharacter.
func (s *service) Update(weaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error) {
	updatedWeaponRelationship, err := s.weaponXCharacterDataRepo.Update(weaponXCharacterData)
    if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
    updatedWeaponRelationship.Weapon, err = s.weaponService.GetById(weaponXCharacterData.Weapon.Weapon_Id)
    if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
    return updatedWeaponRelationship, nil
}

func (s *service) UpdateOwnership(weaponXCharacterData domain.WeaponXCharacterData) (error) {
	err := s.weaponXCharacterDataRepo.UpdateOwnership(weaponXCharacterData)
    if err!= nil {
        return err
    }

    return nil
}

func NewWeaponXCharacterDataService(weaponXCharacterDataRepo RepositoryWeaponXCharacterData, weaponService weapon.ServiceWeapon) ServiceWeaponXCharacterData {
	return &service{weaponXCharacterDataRepo: weaponXCharacterDataRepo, weaponService: weaponService}
}
