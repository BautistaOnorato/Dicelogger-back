package session

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type sessionMySqlRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionMySqlRepository{db: db}
}

func (r *sessionMySqlRepository) Create(session domain.Session) (domain.Session, error) {
	statement, err := r.db.Prepare(QueryCreateSession)
	if err != nil {
		fmt.Println(err)
		return domain.Session{}, ErrPrepareStatement
	}
	
	defer statement.Close()
	result, err := statement.Exec(
		session.Start,
		session.End,
		session.Description,
		session.CampaignId,
		session.CurrentEnviroment,
	)
	if err != nil {
		fmt.Println(err)
		return domain.Session{}, err
	}
	
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return domain.Session{}, ErrGettingLastInsertId
	}
	session.SessionId = int(lastId)
	
	return session, nil
}


func (r *sessionMySqlRepository) GetAll() ([]domain.Session, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []domain.Session
	for rows.Next() {
		var session domain.Session
		if err := rows.Scan(&session.SessionId, &session.Start, &session.End, &session.Description, &session.CampaignId, &session.CurrentEnviroment); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (r *sessionMySqlRepository) GetById(id int) (domain.Session, error) {
	var session domain.Session
	err := r.db.QueryRow(QueryGetById, id).Scan(&session.SessionId, &session.Start, &session.End, &session.Description, &session.CampaignId, &session.CurrentEnviroment)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Session{}, errors.New("session not found")
		}
		return domain.Session{}, err
	}
	return session, nil
}

func (r *sessionMySqlRepository) GetByCampaignId(id int) ([]domain.Session, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []domain.Session
	for rows.Next() {
		var session domain.Session
		if err := rows.Scan(&session.SessionId, &session.Start, &session.End, &session.Description, &session.CampaignId, &session.CurrentEnviroment); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (r *sessionMySqlRepository) Update(session domain.Session, id int) (domain.Session, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		log.Println(err)
		return domain.Session{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(&session.Start, &session.End, &session.Description, &session.CampaignId, &session.CurrentEnviroment, id)
	if err != nil {
		return domain.Session{}, err
	}

	session.SessionId = id
	return session, nil
}

func (r *sessionMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
