package classXspell

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("Row was not found")
)

type ClassXSpellSqlRepository struct {
	db *sql.DB
}

func NewClassXSpellRepository(db *sql.DB) RepositoryClassXSpell {
	return &ClassXSpellSqlRepository{db: db}
}

func (r *ClassXSpellSqlRepository) Create(classXspell domain.ClassXSpell) (domain.ClassXSpell, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.ClassXSpell{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		classXspell.ClassId,
		classXspell.SpellId,
	)
	if err != nil {
		return domain.ClassXSpell{}, err
	}

	return classXspell, nil
}
func (r *ClassXSpellSqlRepository) Delete(classXspell domain.ClassXSpell) error {
	result, err := r.db.Exec(QueryDelete, classXspell.ClassId, classXspell.SpellId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}
