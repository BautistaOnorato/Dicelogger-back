package backgroundXproficiency

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrLastIndex        = errors.New("error getting last index")
)

type backgroundXProficiency struct {
	db *sql.DB
}

func NewBackgroundXProficiencyRepository(db *sql.DB) BackgroundXProficiencyRepository {
	return &backgroundXProficiency{db: db}
}

func (r *backgroundXProficiency) Create(backgroundXProficiency domain.BackgroundXProficiency) (domain.BackgroundXProficiency, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.BackgroundXProficiency{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		backgroundXProficiency.BackgroundID,
		backgroundXProficiency.ProficiencyID,
	)
	if err != nil {
		return domain.BackgroundXProficiency{}, err
	}
	return backgroundXProficiency, nil
}

func (r *backgroundXProficiency) Delete(backgroundXProficiency domain.BackgroundXProficiency) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		backgroundXProficiency.BackgroundID,
		backgroundXProficiency.ProficiencyID,
	)
	if err != nil {
		return err
	}
	return nil
}
