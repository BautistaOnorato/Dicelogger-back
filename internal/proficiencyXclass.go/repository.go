package proficiencyXclass

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
)

type proficiencyXClassSqlRepository struct {
	db *sql.DB
}

func NewProficiencyXClassRepository(db *sql.DB) ProficiencyXClassRepository {
	return &proficiencyXClassSqlRepository{db: db}
}

func (r *proficiencyXClassSqlRepository) Create(proficiencyXClass domain.ProficiencyXClass) (domain.ProficiencyXClass, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.ProficiencyXClass{}, ErrPrepareStatement
	}
	_, err = statement.Exec(
		proficiencyXClass.ClassId,
		proficiencyXClass.ProficiencyId,
	)
	if err != nil {
		return domain.ProficiencyXClass{}, err
	}

	return proficiencyXClass, nil
}

func (r *proficiencyXClassSqlRepository) Delete(proficiencyXClass domain.ProficiencyXClass) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	_, err = statement.Exec(
		proficiencyXClass.ClassId,
		proficiencyXClass.ProficiencyId,
	)
	if err != nil {
		return err
	}

	return nil
}
