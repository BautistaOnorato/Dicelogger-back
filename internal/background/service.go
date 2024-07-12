package background

import (
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type backgroundService struct {
	backgroundRepo BackgroundRepository
}

func NewBackgroundService(backgroundRepo BackgroundRepository) BackgroundService {
	return &backgroundService{backgroundRepo: backgroundRepo}
}

func (s *backgroundService) CreateBackground(backgroundDto dto.CreateBackgroundDto) (domain.Background, error) {
	backgroundDomain := domain.Background{
		BackgroundID:      backgroundDto.BackgroundID,
		Name:              backgroundDto.Name,
		Languages:         backgroundDto.Languages,
		PersonalityTraits: backgroundDto.PersonalityTraits,
		Ideals:            backgroundDto.Ideals,
		Bond:              backgroundDto.Bond,
		Flaws:             backgroundDto.Flaws,
		Trait:             backgroundDto.Trait,
		ToolProficiencies: backgroundDto.ToolProficiencies,
	}

	createdBackground, err := s.backgroundRepo.Create(backgroundDomain)
	if err != nil {
		return domain.Background{}, err
	}

	return createdBackground, nil
}

func (s *backgroundService) GetAllBackgrounds() ([]domain.Background, error) {
	backgrounds, err := s.backgroundRepo.GetAllBackgrounds()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return backgrounds, nil
}

func (s *backgroundService) GetBackgroundByID(id int) (domain.Background, error) {
	background, err := s.backgroundRepo.GetBackgroundById(id)
	if err != nil {
		return domain.Background{}, err
	}
	return background, nil
}

func (s *backgroundService) UpdateBackground(backgroundDto dto.CreateBackgroundDto, id int) (domain.Background, error) {
	backgroundDomain := domain.Background{
		BackgroundID:      backgroundDto.BackgroundID,
		Name:              backgroundDto.Name,
		Languages:         backgroundDto.Languages,
		PersonalityTraits: backgroundDto.PersonalityTraits,
		Ideals:            backgroundDto.Ideals,
		Bond:              backgroundDto.Bond,
		Flaws:             backgroundDto.Flaws,
		Trait:             backgroundDto.Trait,
		ToolProficiencies: backgroundDto.ToolProficiencies,
	}

	updatedBackground, err := s.backgroundRepo.UpdateBackground(backgroundDomain, id)
	if err != nil {
		fmt.Println(err)
		return domain.Background{}, err
	}

	return updatedBackground, nil
}

func (s *backgroundService) DeleteBackground(id int) error {
	return s.backgroundRepo.DeleteBackground(id)
}
