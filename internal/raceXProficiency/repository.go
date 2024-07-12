package raceXproficiency

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrLastIndex        = errors.New("error getting last index")
)

type raceXProficiencyMySqlRepository struct {
	db *sql.DB
}

func NewRaceXProficiencyRepository(db *sql.DB) RaceXProficiencyRepository {
	return &raceXProficiencyMySqlRepository{db: db}
}

func (r *raceXProficiencyMySqlRepository) Create(raceXProficiency domain.RaceXProficiency) (domain.RaceXProficiency, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.RaceXProficiency{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		raceXProficiency.RaceId,
		raceXProficiency.ProficiencyId,
	)
	if err != nil {
		return domain.RaceXProficiency{}, err
	}

	return raceXProficiency, nil
}

func (r *raceXProficiencyMySqlRepository) Delete(raceXProficiency domain.RaceXProficiency) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		raceXProficiency.RaceId,
		raceXProficiency.ProficiencyId,
	)
	if err != nil {
		return err
	}

	return nil
}
