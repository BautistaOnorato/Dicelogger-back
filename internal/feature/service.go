package feature

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo FeatureRepository
}

func NewFeatureService(repo FeatureRepository) FeatureService {
	return &service{repo: repo}
}

func (s *service) CreateFeature(featureDto dto.CreateFeatureDto) (domain.Feature, error) {
	// featureDomain := domain.Feature{
	// 	Name:        featureDto.Name,
	// 	Description: featureDto.Description,
	// }

	createdFeature, err := s.repo.Create(featureDto)
	if err != nil {
		return domain.Feature{}, err
	}

	return createdFeature, nil
}

func (s *service) GetAllFeatures() ([]domain.Feature, error) {
	features, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return features, nil
}

func (s *service) GetAllFeaturesByCharacterId(characterId int) (dto.FeatureFullResponseDto, error) {
	features, err := s.repo.GetAllByCharacterId(characterId)
	if err != nil {
		return dto.FeatureFullResponseDto{}, err
	}

	var response dto.FeatureFullResponseDto = dto.FeatureFullResponseDto{
		CharacterId: characterId,
		Features:    features,
	}

	return response, nil
}

func (s *service) GetFeatureById(id int) (domain.Feature, error) {
	feature, err := s.repo.GetById(id)
	if err != nil {
		return domain.Feature{}, err
	}

	return feature, nil
}

func (s *service) UpdateFeature(featureDto dto.CreateFeatureDto, id int) (domain.Feature, error) {
	featureDomain := domain.Feature{
		Name:        featureDto.Name,
		Description: featureDto.Description,
	}

	updatedFeature, err := s.repo.Update(featureDomain, id)
	if err != nil {
		return domain.Feature{}, err
	}

	return updatedFeature, nil
}

func (s *service) DeleteFeature(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
