package raceXproficiency

import "github.com/proyecto-dnd/backend/internal/domain"

type RaceXProficiencyRepository interface {
	Create(raceXProficiency domain.RaceXProficiency) (domain.RaceXProficiency, error)
	Delete(raceXProficiency domain.RaceXProficiency) error
}

type RaceXProficiencyService interface {
	Create(raceXProficiency domain.RaceXProficiency) (domain.RaceXProficiency, error)
	Delete(raceXProficiency domain.RaceXProficiency) error
}
