package character_feature

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo CharacterFeatureRepository
}

func NewCharacterFeatureService(repo CharacterFeatureRepository) CharacterFeatureService {
	return &service{repo: repo}
}

// DeleteByCharacterDataId implements CharacterFeatureService.
func (s *service) DeleteByCharacterDataId(idCharacter int) error {
	err := s.repo.DeleteByCharacterDataId(idCharacter)
	if err != nil {
		return err
	}

	return nil
}


func (s *service) CreateCharacterFeature(characterFeatureDto dto.CreateCharacterFeatureDto) (domain.CharacterFeature, error) {
	characterFeatureDomain := domain.CharacterFeature{
		CharacterId: characterFeatureDto.CharacterId,
		FeatureId:   characterFeatureDto.FeatureId,
	}

	createdCharacterFeature, err := s.repo.Create(characterFeatureDomain)
	if err != nil {
		return domain.CharacterFeature{}, err
	}

	return createdCharacterFeature, nil
}

func (s *service) GetAllCharacterFeatures() ([]domain.CharacterFeature, error) {
	characterFeatures, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return characterFeatures, nil
}

func (s *service) GetCharacterFeatureByFeatureId(id int) ([]domain.CharacterFeature, error) {
	characterFeatures, err := s.repo.GetByFeatureId(id)
	if err != nil {
		return nil, err
	}

	return characterFeatures, nil
}

func (s *service) GetCharacterFeatureByCharacterId(id int) ([]domain.CharacterFeature, error) {
	characterFeatures, err := s.repo.GetByCharacterId(id)
	if err != nil {
		return nil, err
	}

	return characterFeatures, nil
}

func (s *service) DeleteCharacterFeature(idFeature int, idCharacter int) error {
	err := s.repo.Delete(idFeature, idCharacter)
	if err != nil {
		return err
	}

	return nil
}
