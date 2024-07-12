package race

import (
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type raceService struct {
	raceRepo RaceRepository
}

func NewRaceService(raceRepo RaceRepository) RaceService {
	return &raceService{raceRepo: raceRepo}
}

func (s *raceService) CreateRace(raceDto dto.CreateRaceDto) (domain.Race, error) {
	raceDomain := domain.Race{
		Name:        raceDto.Name,
		Description: raceDto.Description,
		Speed:       raceDto.Speed,
		Str:         raceDto.Str,
		Dex:         raceDto.Dex,
		Int:        raceDto.Int,
		Con:         raceDto.Con,
		Wiz:         raceDto.Wiz,
		Cha:         raceDto.Cha,
	}

	createdRace, err := s.raceRepo.Create(raceDomain)
	if err != nil {
		return domain.Race{}, err
	}

	return createdRace, nil
}

func (s *raceService) GetAllRaces() ([]domain.Race, error) {
	races, err := s.raceRepo.GetAllRaces()
	if err != nil {
		return nil, err
	}

	return races, nil
}

func (s *raceService) GetRaceByID(id int) (domain.Race, error) {
	race, err := s.raceRepo.GetRaceById(id)
	if err != nil {
		return domain.Race{}, err
	}

	return race, nil
}

func (s *raceService) UpdateRace(raceDto dto.CreateRaceDto, id int) (domain.Race, error) {
	raceDomain := domain.Race{
		Name:        raceDto.Name,
		Description: raceDto.Description,
		Speed:       raceDto.Speed,
		Str:         raceDto.Str,
		Dex:         raceDto.Dex,
		Int:        raceDto.Int,
		Con:         raceDto.Con,
		Wiz:         raceDto.Wiz,
		Cha:         raceDto.Cha,
	}

	updatedRace, err := s.raceRepo.UpdateRace(raceDomain, id)
	if err != nil {
		fmt.Println(err)
		return domain.Race{}, err
	}

	return updatedRace, nil
}

func (s *raceService) DeleteRace(id int) error {
	return s.raceRepo.DeleteRace(id)
}
