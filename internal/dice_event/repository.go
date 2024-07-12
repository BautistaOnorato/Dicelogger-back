package dice_event

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrLastInsertId     = errors.New("error getting last insert id")
)

type repository struct {
	db *sql.DB
}

func NewDiceEventRepository(db *sql.DB) DiceEventRepository {
	return &repository{db: db}
}

// DeleteByProtagonistId implements DiceEventRepository.
func (r *repository) DeleteByProtagonistId(id int) error {
	statement, err := r.db.Prepare(QueryDeleteByProtagonistId)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}


// GetBySessionId implements DiceEventRepository.
func (r *repository) GetBySessionId(sessionid int) ([]domain.DiceEvent, error) {
	rows, err := r.db.Query(QueryGetBySessionId, sessionid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var diceEvents []domain.DiceEvent
	for rows.Next() {
		var diceEvent domain.DiceEvent
		if err := rows.Scan(
			&diceEvent.DiceEventId,
			&diceEvent.Stat,
			&diceEvent.Difficulty,
			&diceEvent.DiceRolled,
			&diceEvent.DiceResult,
			&diceEvent.EventProtagonist,
			&diceEvent.Description,
			&diceEvent.SessionId,
			&diceEvent.TimeStamp,
		); err != nil {
			return nil, err
		}
		diceEvents = append(diceEvents, diceEvent)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return diceEvents, nil
}

func (r *repository) Create(diceEvent domain.DiceEvent) (domain.DiceEvent, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.DiceEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		diceEvent.DiceEventId,
		diceEvent.Stat,
		diceEvent.Difficulty,
		diceEvent.DiceRolled,
		diceEvent.DiceResult,
		diceEvent.EventProtagonist,
		diceEvent.Description,
		diceEvent.SessionId,
		diceEvent.TimeStamp,
	)
	if err != nil {
		return domain.DiceEvent{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.DiceEvent{}, ErrLastInsertId
	}
	diceEvent.DiceEventId = int(lastId)
	return domain.DiceEvent{}, nil
}

func (r *repository) GetAll() ([]domain.DiceEvent, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var diceEvents []domain.DiceEvent
	for rows.Next() {
		var diceEvent domain.DiceEvent
		if err := rows.Scan(
			&diceEvent.DiceEventId,
			&diceEvent.Stat,
			&diceEvent.Difficulty,
			&diceEvent.DiceRolled,
			&diceEvent.DiceResult,
			&diceEvent.EventProtagonist,
			&diceEvent.Description,
			&diceEvent.SessionId,
			&diceEvent.TimeStamp,
		); err != nil {
			return nil, err
		}
		diceEvents = append(diceEvents, diceEvent)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return diceEvents, nil
}

func (r *repository) GetById(id int) (domain.DiceEvent, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var diceEvent domain.DiceEvent
	err := row.Scan(
		&diceEvent.DiceEventId,
		&diceEvent.Stat,
		&diceEvent.Difficulty,
		&diceEvent.DiceRolled,
		&diceEvent.DiceResult,
		&diceEvent.EventProtagonist,
		&diceEvent.Description,
		&diceEvent.SessionId,
		&diceEvent.TimeStamp,
	)
	if err != nil {
		return domain.DiceEvent{}, err
	}
	return diceEvent, nil
}

func (r *repository) Update(diceEvent domain.DiceEvent, id int) (domain.DiceEvent, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.DiceEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		diceEvent.Stat,
		diceEvent.Difficulty,
		diceEvent.DiceRolled,
		diceEvent.DiceResult,
		diceEvent.EventProtagonist,
		diceEvent.Description,
		diceEvent.SessionId,
		id,
	)
	if err != nil {
		return domain.DiceEvent{}, err
	}
	return diceEvent, nil
}

func (r *repository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
