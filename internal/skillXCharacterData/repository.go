package skillxcharacterdata

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("skills not found")
)

type skillxCharacterDataSqlRepository struct {
	db *sql.DB
}

func NewSkillxCharacterDataRepository(db *sql.DB) RepositorySkillXCharacter {
	return &skillxCharacterDataSqlRepository{db}
}

// DeleteByCharacterDataId implements RepositorySkillXCharacter.
func (r *skillxCharacterDataSqlRepository) DeleteByCharacterDataId(id int) error {
	result, err := r.db.Exec(QueryDeleteByCharacterId, id)
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


// Create implements RepositorySkillXCharacterData.
func (r *skillxCharacterDataSqlRepository) Create(skillXCharacterData domain.SkillXCharacterData) (domain.SkillXCharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateSkillXCharacter)
	if err != nil {
		return domain.SkillXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		skillXCharacterData.SkillID,
		skillXCharacterData.CharacterID,
	)

	if err != nil {
		return domain.SkillXCharacterData{}, err
	}

	return skillXCharacterData, nil
}

// Delete implements RepositorySkillXCharacterData.
func (r *skillxCharacterDataSqlRepository) Delete(skillXCharacterData domain.SkillXCharacterData) error {
	result, err := r.db.Exec(QueryDeleteSkillXCharacter, skillXCharacterData.SkillID, skillXCharacterData.CharacterID, nil)
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
