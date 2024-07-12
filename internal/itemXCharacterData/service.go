package itemxcharacterdata

import (
	"fmt"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/item"
)

type service struct {
	itemXCharacterDataRepo RepositoryItemXCharacterData
	itemService            item.ServiceItem
}

// UpdateOwnership implements ServiceItemXCharacterData.

func NewItemXCharacterDataService(itemXCharacterDataRepo RepositoryItemXCharacterData, itemService item.ServiceItem) ServiceItemXCharacterData {
    return &service{itemXCharacterDataRepo: itemXCharacterDataRepo, itemService: itemService}
}

// Create implements ServiceItemXTableCharacter.
func (s *service) Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
    newItemRelationship, err := s.itemXCharacterDataRepo.Create(itemXCharacterData)
	if err != nil {
        return domain.ItemXCharacterData{}, err
	}
	newItemRelationship.Item, err = s.itemService.GetById(itemXCharacterData.Item.Item_Id)
	if err != nil {
        return domain.ItemXCharacterData{}, err
	}
	return newItemRelationship, nil
}

// Delete implements ServiceItemXTableCharacter.
func (s *service) Delete(id int) error {
    err := s.itemXCharacterDataRepo.Delete(id)
	if err != nil {
        return err
	}
	return nil
}

// DeleteByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) DeleteByCharacterDataId(id int) error {
    err := s.itemXCharacterDataRepo.DeleteByCharacterDataId(id)
	if err != nil {
        return err
	}
	return nil
}

// GetAll implements ServiceItemXTableCharacter.
func (s *service) GetAll() ([]domain.ItemXCharacterData, error) {
    itemRelationships, err := s.itemXCharacterDataRepo.GetAll()
	if err != nil {
        return []domain.ItemXCharacterData{}, err
	}
	return itemRelationships, nil
}

// GetByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) GetByCharacterDataId(id int) ([]domain.ItemXCharacterData, error) {
    itemRelationships, err := s.itemXCharacterDataRepo.GetByCharacterDataId(id)
	if err != nil {
        return []domain.ItemXCharacterData{}, err
	}
	return itemRelationships, nil
}

// GetById implements ServiceItemXTableCharacter.
func (s *service) GetById(id int) (domain.ItemXCharacterData, error) {
    itemRelationship, err := s.itemXCharacterDataRepo.GetById(id)
	if err != nil {
        return domain.ItemXCharacterData{}, err
	}
	return itemRelationship, nil
}

// Update implements ServiceItemXTableCharacter.
func (s *service) Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
    newItemRelationship, err := s.itemXCharacterDataRepo.Update(itemXCharacterData)
	if err != nil {
        return domain.ItemXCharacterData{}, err
	}
	newItemRelationship.Item, err = s.itemService.GetById(itemXCharacterData.Item.Item_Id)
	if err != nil {
        return domain.ItemXCharacterData{}, err
	}
	return newItemRelationship, nil
}

func (s *service) UpdateOwnership(itemXCharacterData domain.ItemXCharacterData) error {
	fmt.Println("waos1")
    err := s.itemXCharacterDataRepo.UpdateOwnership(itemXCharacterData)
	if err != nil {
        return err
	}
	
	return nil
}
