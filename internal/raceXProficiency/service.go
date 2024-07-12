package raceXproficiency

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	repository RaceXProficiencyRepository
}

func NewRaceXProficiencyService(repository RaceXProficiencyRepository) *service {
	return &service{repository: repository}
}

func (s *service) Create(raceXProficiency domain.RaceXProficiency) (domain.RaceXProficiency, error) {
	newRacexRepository, err := s.repository.Create(raceXProficiency)
	if err != nil {
		return domain.RaceXProficiency{}, err
	}

	return newRacexRepository, nil
}

func (s *service) Delete(raceXProficiency domain.RaceXProficiency) error {
	return s.repository.Delete(raceXProficiency)
}
