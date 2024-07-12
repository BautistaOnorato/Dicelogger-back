package armorXCharacterData

import (
	"github.com/proyecto-dnd/backend/internal/armor"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	armorXcharacterRepo RepositoryArmorXCharacterData
	armorService armor.ArmorService
}

func NewServiceArmorXCharacterData(armorXcharacterRepo RepositoryArmorXCharacterData, armorService armor.ArmorService) ServiceArmorXCharacterData {
	return &service{armorXcharacterRepo: armorXcharacterRepo, armorService: armorService}
}

// CreateArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error) {
	newArmorRelationship, err := s.armorXcharacterRepo.CreateArmorXCharacterData(data)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	newArmorRelationship.Armor, err = s.armorService.GetArmorByID(data.Armor.ArmorId)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	return newArmorRelationship, nil
}

// DeleteArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) DeleteArmorXCharacterData(id int) error {
	err := s.armorXcharacterRepo.DeleteArmorXCharacterData(id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByCharacterDataIdArmor implements ServiceArmorXCharacterData.
func (s *service) DeleteByCharacterDataIdArmor(id int) error {
	err := s.armorXcharacterRepo.DeleteByCharacterDataIdArmor(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error) {
	armorRelationships, err := s.armorXcharacterRepo.GetAllArmorXCharacterData()
	if err != nil {
		return []domain.ArmorXCharacterData{}, err
	}
	return armorRelationships, nil
}

// GetByIdArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) GetByIdArmorXCharacterData(id int) (domain.ArmorXCharacterData, error) {
	armorRelationship, err := s.armorXcharacterRepo.GetByIdArmorXCharacterData(id)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	return armorRelationship, nil
}

// GetByCharacterDataIdArmor implements ServiceArmorXCharacterData.
func (s *service) GetByCharacterDataIdArmor(id int) ([]domain.ArmorXCharacterData, error) {
	armorRelationships, err := s.armorXcharacterRepo.GetByCharacterDataIdArmor(id)
	if err != nil {
		return []domain.ArmorXCharacterData{}, err
	}
	return armorRelationships, nil
}

// UpdateArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error) {
	updatedArmorRelationship, err := s.armorXcharacterRepo.UpdateArmorXCharacterData(data)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	updatedArmorRelationship.Armor, err = s.armorService.GetArmorByID(data.Armor.ArmorId)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	return updatedArmorRelationship, nil
}

func (s *service) UpdateOwnership(data domain.ArmorXCharacterData) (error) {
	err := s.armorXcharacterRepo.UpdateOwnership(data)
	if err != nil {
		return err
	}
	return nil
}
