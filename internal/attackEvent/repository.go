package attackEvent

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type attackEventMySqlRepository struct {
	db *sql.DB
}

func NewAttackEventRepository(db *sql.DB) AttackEventRepository {
	return &attackEventMySqlRepository{db: db}
}

func (r *attackEventMySqlRepository) Create(event domain.AttackEvent) (domain.AttackEvent, error) {
	statement, err := r.db.Prepare(QueryCreateAttackEvent)
	if err != nil {
		return domain.AttackEvent{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		event.Type,
		event.Environment,
		event.Session_id,
		event.EventProtagonistId,
		event.EventResolution,
		event.Weapon,
		event.Spell,
		event.DmgType,
		event.Description,
		event.TimeStamp,
	)
	if err != nil {
		return domain.AttackEvent{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.AttackEvent{}, ErrGettingLastInsertId
	}
	event.AttackEventId = int(lastId)

	return event, nil
}

func (r *attackEventMySqlRepository) GetAll() ([]domain.AttackEvent, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.AttackEvent
	for rows.Next() {
		var event domain.AttackEvent
		if err := rows.Scan(&event.AttackEventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.EventResolution, &event.Weapon, &event.Spell, &event.DmgType, &event.Description, &event.TimeStamp); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *attackEventMySqlRepository) GetById(id int) (dto.RepositoryResponseAttackEvent, error) {
	var event dto.RepositoryResponseAttackEvent
	err := r.db.QueryRow(QueryGetById, id).Scan(&event.AttackEventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.EventResolution, &event.Weapon, &event.Spell, &event.DmgType, &event.Description, &event.TimeStamp, &event.SessionSessionId, &event.Start, &event.End, &event.SessionDescription, &event.SessionCampaignId, &event.SessionCurrentEnviroment, &event.CharacterId, &event.CharacterUserId, &event.CharacterCampaignId, &event.CharacterName, &event.ClassName, &event.RaceName, &event.Level, &event.HitPoints)
	if err != nil {
		fmt.Println(err)
		return dto.RepositoryResponseAttackEvent{}, err
	}
	return event, nil
}

func (r *attackEventMySqlRepository) GetBySessionId(session_id int) ([]dto.RepositoryResponseAttackEvent, error) {
	rows, err := r.db.Query(QueryGetBySessionId, session_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.RepositoryResponseAttackEvent
	for rows.Next() {
		var event dto.RepositoryResponseAttackEvent
		if err := rows.Scan(&event.AttackEventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.EventResolution, &event.Weapon, &event.Spell, &event.DmgType, &event.Description, &event.TimeStamp, &event.SessionSessionId, &event.Start, &event.End, &event.SessionDescription, &event.SessionCampaignId, &event.SessionCurrentEnviroment, &event.CharacterId, &event.CharacterUserId, &event.CharacterCampaignId, &event.CharacterName, &event.ClassName, &event.RaceName, &event.Level, &event.HitPoints); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *attackEventMySqlRepository) GetByProtagonistId(protagonist_id int) ([]dto.RepositoryResponseAttackEvent, error) {
	rows, err := r.db.Query(QueryGetBySessionId, protagonist_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.RepositoryResponseAttackEvent
	for rows.Next() {
		var event dto.RepositoryResponseAttackEvent
		if err := rows.Scan(&event.AttackEventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.EventResolution, &event.Weapon, &event.Spell, &event.DmgType, &event.Description, &event.TimeStamp, &event.SessionSessionId, &event.Start, &event.End, &event.SessionDescription, &event.SessionCampaignId, &event.SessionCurrentEnviroment, &event.CharacterId, &event.CharacterUserId, &event.CharacterCampaignId, &event.CharacterName, &event.ClassName, &event.RaceName, &event.Level, &event.HitPoints); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *attackEventMySqlRepository) GetByAffectedId(affected_id int) ([]dto.RepositoryResponseAttackEvent, error) {
	rows, err := r.db.Query(QueryGetByAffectedId, affected_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.RepositoryResponseAttackEvent
	for rows.Next() {
		var event dto.RepositoryResponseAttackEvent
		if err := rows.Scan(&event.AttackEventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.EventResolution, &event.Weapon, &event.Spell, &event.DmgType, &event.Description, &event.TimeStamp, &event.SessionSessionId, &event.Start, &event.End, &event.SessionDescription, &event.SessionCampaignId, &event.SessionCurrentEnviroment, &event.CharacterId, &event.CharacterUserId, &event.CharacterCampaignId, &event.CharacterName, &event.ClassName, &event.RaceName, &event.Level, &event.HitPoints); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *attackEventMySqlRepository) GetByProtagonistIdAndAffectedId(protagonist_id int, affected_id int) ([]dto.RepositoryResponseAttackEvent, error) {
	rows, err := r.db.Query(QueryGetByProtagonistIdAndAffectedId, protagonist_id, affected_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.RepositoryResponseAttackEvent
	for rows.Next() {
		var event dto.RepositoryResponseAttackEvent
		if err := rows.Scan(&event.AttackEventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.EventResolution, &event.DmgType, &event.Weapon, &event.Spell, &event.Description, &event.TimeStamp, &event.SessionSessionId, &event.Start, &event.End, &event.SessionDescription, &event.SessionCampaignId, &event.SessionCurrentEnviroment, &event.CharacterId, &event.CharacterUserId, &event.CharacterCampaignId, &event.CharacterName, &event.RaceName, &event.ClassName, &event.Level, &event.HitPoints); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r* attackEventMySqlRepository) GetCharacterDataByAttackEventId(id int) ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetCharacterDataByAttackEventId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []dto.CharacterCardDto
	for rows.Next() {
		var character dto.CharacterCardDto
		if err := rows.Scan(&character.CharacterId, &character.UserId, &character.CampaignID, &character.ImageUrl, &character.Name, &character.Race, &character.Name, &character.Level, &character.HitPoints); err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	return characters, nil
}

func (r *attackEventMySqlRepository) Update(event domain.AttackEvent, id int) (domain.AttackEvent, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.AttackEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		event.Type,
		event.Environment,
		event.Session_id,
		event.EventProtagonistId,
		event.EventResolution,
		event.Weapon,
		event.Spell,
		event.DmgType,
		event.Description,
		event.TimeStamp,
		id,
	)
	if err != nil {
		return domain.AttackEvent{}, err
	}
	event.AttackEventId = id
	return event, nil
}

func (r *attackEventMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
