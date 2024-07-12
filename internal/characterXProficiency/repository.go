package characterXproficiency

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrLastIndex        = errors.New("error getting last index")
)

type characterXProficiencyRepository struct {
	db *sql.DB
}

func NewCharacterXProficiencyRepository(db *sql.DB) CharacterXProficiencyRepository {
	return &characterXProficiencyRepository{db: db}
}

// DeleteByCharacterDataId implements CharacterXProficiencyRepository.
func (r *characterXProficiencyRepository) DeleteByCharacterDataId(id int) error {
	statement, err := r.db.Prepare(QueryDeleteByCharacterId)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}


func (r *characterXProficiencyRepository) Create(characterXProficiency domain.CharacterXProficiency) (domain.CharacterXProficiency, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.CharacterXProficiency{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		characterXProficiency.CharacterId,
		characterXProficiency.ProficiencyId,
	)
	if err != nil {
		return domain.CharacterXProficiency{}, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.CharacterXProficiency{}, ErrLastIndex
	}
	var createdCharacterProficiency = domain.CharacterXProficiency{
		CharacterProficiencyId: int(lastId),
		CharacterId:            characterXProficiency.CharacterId,
		ProficiencyId:          characterXProficiency.ProficiencyId,
	}

	return createdCharacterProficiency, nil
}

func (r *characterXProficiencyRepository) Delete(characterXProficiencyId int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(characterXProficiencyId)
	if err != nil {
		return err
	}

	return nil
}
