package tradeevent

import (
	"database/sql"
	"errors"
	"time"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrNotFound = errors.New("event not found")
)

type tradeEventMysqlRepository struct {
	db *sql.DB
}

func NewTradeEventMySqlRepository(db *sql.DB) RepositoryTradeEvent {
	return &tradeEventMysqlRepository{db}
}

// Create implements RepositoryTradeEvent.
func (t *tradeEventMysqlRepository) Create(tradeEvent domain.TradeEvent) (domain.TradeEvent, error) {
	statement, err := t.db.Prepare(QueryInsert)
	if err != nil {
		return domain.TradeEvent{}, err
	}

	timestamp := time.Now()

	result, err := statement.Exec(
		tradeEvent.Session_Id,
		tradeEvent.Sender,
		tradeEvent.Receiver,
		tradeEvent.Description,
		timestamp,
	)
	if err != nil {
		return domain.TradeEvent{}, err
	}
	lastInsert, err := result.LastInsertId()
	if err != nil {
		return domain.TradeEvent{}, err
	}
	tradeEvent.Timestamp = &timestamp
	tradeEvent.TradeEvent_Id = int(lastInsert)
	return tradeEvent, nil
}

// Delete implements RepositoryTradeEvent.
func (t *tradeEventMysqlRepository) Delete(id int) error {
	result, err := t.db.Exec(QueryDelete, id)
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

// GetByReceiver implements RepositoryTradeEvent.
func (t *tradeEventMysqlRepository) GetByReceiver(receiver int) ([]domain.TradeEvent, error) {
	rows, err := t.db.Query(QueryGetByReceiver, receiver)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	tradeEvents := []domain.TradeEvent{}
	for rows.Next() {
		tradeEvent := domain.TradeEvent{}
		err := rows.Scan(
			&tradeEvent.TradeEvent_Id,
			&tradeEvent.Session_Id,
			&tradeEvent.Sender,
			&tradeEvent.Receiver,
			&tradeEvent.Description,
			&tradeEvent.Timestamp,
		)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
		tradeEvents = append(tradeEvents, tradeEvent)
	}
	if err := rows.Err(); err != nil {
		return []domain.TradeEvent{}, err
	}
	return tradeEvents, nil
}

// GetBySender implements RepositoryTradeEvent.
func (t *tradeEventMysqlRepository) GetBySender(sender int) ([]domain.TradeEvent, error) {
	rows, err := t.db.Query(QueryGetBySender, sender)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	tradeEvents := []domain.TradeEvent{}
	for rows.Next() {
		tradeEvent := domain.TradeEvent{}
		err := rows.Scan(
			&tradeEvent.TradeEvent_Id,
			&tradeEvent.Session_Id,
			&tradeEvent.Sender,
			&tradeEvent.Receiver,
			&tradeEvent.Description,
			&tradeEvent.Timestamp,
		)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
		tradeEvents = append(tradeEvents, tradeEvent)
	}
	if err := rows.Err(); err != nil {
		return []domain.TradeEvent{}, err
	}
	return tradeEvents, nil
}

// GetBySenderOrReciever implements RepositoryTradeEvent.
func (t *tradeEventMysqlRepository) GetBySenderOrReciever(id int) ([]domain.TradeEvent, error) {
	rows, err := t.db.Query(QueryGetBySenderOrReciever, id, id)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	tradeEvents := []domain.TradeEvent{}
	for rows.Next() {
		tradeEvent := domain.TradeEvent{}
		err := rows.Scan(
			&tradeEvent.TradeEvent_Id,
			&tradeEvent.Session_Id,
			&tradeEvent.Sender,
			&tradeEvent.Receiver,
			&tradeEvent.Description,
			&tradeEvent.Timestamp,
		)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
		tradeEvents = append(tradeEvents, tradeEvent)
	}
	if err := rows.Err(); err != nil {
		return []domain.TradeEvent{}, err
	}
	return tradeEvents, nil
}

// GetBySessionId implements RepositoryTradeEvent.
func (t *tradeEventMysqlRepository) GetBySessionId(sessionId int) ([]domain.TradeEvent, error) {
	rows, err := t.db.Query(QueryGetBySessionId, sessionId)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	tradeEvents := []domain.TradeEvent{}
	for rows.Next() {
		tradeEvent := domain.TradeEvent{}
		err := rows.Scan(
			&tradeEvent.TradeEvent_Id,
			&tradeEvent.Session_Id,
			&tradeEvent.Sender,
			&tradeEvent.Receiver,
			&tradeEvent.Description,
			&tradeEvent.Timestamp,
		)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
		tradeEvents = append(tradeEvents, tradeEvent)
	}
	if err := rows.Err(); err != nil {
		return []domain.TradeEvent{}, err
	}
	return tradeEvents, nil
}
