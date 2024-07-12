package characterxattackevent

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("characters not found")
)

type characterXAttackEventRepository struct {
	db *sql.DB
}

func NewCharacterXAttackEventRepository(db *sql.DB) CharacterXAttackEventRepository {
	return &characterXAttackEventRepository{db: db}
}

func (r *characterXAttackEventRepository) GetAll() ([]domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetAll)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	var characterXAttackEvents []domain.CharacterXAttackEvent
	for rows.Next() {
		var characterXAttackEvent domain.CharacterXAttackEvent
		err := rows.Scan(&characterXAttackEvent.CharacterAttackEventId, &characterXAttackEvent.EventId, &characterXAttackEvent.CharacterId, &characterXAttackEvent.Dmg, &characterXAttackEvent.DmgRoll, &characterXAttackEvent.AttackResult, &characterXAttackEvent.AttackRoll, &characterXAttackEvent.ArmorClass)
		if err != nil {
			return nil, err
		}
		characterXAttackEvents = append(characterXAttackEvents, characterXAttackEvent)
	}
	return characterXAttackEvents, nil
}

func (r *characterXAttackEventRepository) GetById(id int) (domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetById)
	if err != nil {
		fmt.Println(err)
		return domain.CharacterXAttackEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	var characterXAttackEvent domain.CharacterXAttackEvent
	err = statement.QueryRow(id).Scan(&characterXAttackEvent.CharacterAttackEventId, &characterXAttackEvent.EventId, &characterXAttackEvent.CharacterId, &characterXAttackEvent.Dmg, &characterXAttackEvent.DmgRoll, &characterXAttackEvent.AttackResult, &characterXAttackEvent.AttackRoll, &characterXAttackEvent.ArmorClass)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}
	return characterXAttackEvent, nil
}

func (r *characterXAttackEventRepository) GetByCharacterId(characterId int) ([]domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetByCharacterId)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(characterId)
	if err != nil {
		return nil, err
	}

	var characterXAttackEvents []domain.CharacterXAttackEvent
	for rows.Next() {
		var characterXAttackEvent domain.CharacterXAttackEvent
		err := rows.Scan(&characterXAttackEvent.CharacterAttackEventId, &characterXAttackEvent.EventId, &characterXAttackEvent.CharacterId, &characterXAttackEvent.Dmg, &characterXAttackEvent.DmgRoll, &characterXAttackEvent.AttackResult, &characterXAttackEvent.AttackRoll, &characterXAttackEvent.ArmorClass)
		if err != nil {
			return nil, err
		}
		characterXAttackEvents = append(characterXAttackEvents, characterXAttackEvent)
	}
	return characterXAttackEvents, nil
}

func (r *characterXAttackEventRepository) GetByEventId(EventId int) ([]domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetByEventId)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(EventId)
	if err != nil {
		return nil, err
	}

	var characterXAttackEvents []domain.CharacterXAttackEvent
	for rows.Next() {
		var characterXAttackEvent domain.CharacterXAttackEvent
		err := rows.Scan(&characterXAttackEvent.CharacterAttackEventId, &characterXAttackEvent.EventId, &characterXAttackEvent.CharacterId, &characterXAttackEvent.Dmg, &characterXAttackEvent.DmgRoll, &characterXAttackEvent.AttackResult, &characterXAttackEvent.AttackRoll, &characterXAttackEvent.ArmorClass)
		if err != nil {
			return nil, err
		}
		characterXAttackEvents = append(characterXAttackEvents, characterXAttackEvent)
	}
	return characterXAttackEvents, nil
}

func (r *characterXAttackEventRepository) Create(characterXAttackEvent domain.CharacterXAttackEvent) (domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.CharacterXAttackEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(characterXAttackEvent.EventId, characterXAttackEvent.CharacterId, characterXAttackEvent.Dmg, characterXAttackEvent.DmgRoll, characterXAttackEvent.AttackResult, characterXAttackEvent.AttackRoll, characterXAttackEvent.ArmorClass)
	if err != nil {
		fmt.Println(err)
		return domain.CharacterXAttackEvent{}, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return domain.CharacterXAttackEvent{}, ErrGettingLastInsertId
	}

	characterXAttackEvent.CharacterAttackEventId = int(lastInsertId)
	return characterXAttackEvent, nil
}

func (r *characterXAttackEventRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDelete, id)
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