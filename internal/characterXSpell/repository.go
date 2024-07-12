package characterXspell

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("spells not found")
)

type CharacterXSpellRepository struct {
	db *sql.DB
}

func NewCharacterXSpellRepository(db *sql.DB) RepositoryCharacterXSpell {
	return &CharacterXSpellRepository{db: db}
}

// DeleteByCharacterDataId implements RepositoryCharacterXSpell.
func (r *CharacterXSpellRepository) DeleteByCharacterDataId(id int) error {
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


func (r *CharacterXSpellRepository) Create(characterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.CharacterXSpell{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		characterXSpell.CharacterId,
		characterXSpell.SpellId,
	)
	if err != nil {
		return domain.CharacterXSpell{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.CharacterXSpell{}, ErrGettingLastInsertId
	}
	var createdCharacterXSpell = domain.CharacterXSpell{
		CharacterSpellId: int(lastId),
		CharacterId:      characterXSpell.CharacterId,
		SpellId:          characterXSpell.SpellId,
	}
	fmt.Println(createdCharacterXSpell)
	return createdCharacterXSpell, nil
}

func (r *CharacterXSpellRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
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

func (r *CharacterXSpellRepository) DeleteParams(characterId int, spellId int) error {
	statement, err := r.db.Prepare(QueryDeleteParams)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(characterId, spellId)
	if err != nil {
		return err
	}
	return nil
}
