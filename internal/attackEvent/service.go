package attackEvent

import (
	"time"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo              AttackEventRepository
}

func NewAttackEventService(repo AttackEventRepository) AttackEventService {
	return &service{repo: repo}
}

// DeleteByProtagonistAndAffectedId implements AttackEventService.
func (s *service) DeleteByProtagonistAndAffectedId(protagonistId int, affectedId int) error {
	attackEvents, err := s.GetEventsByProtagonistIdAndAffectedId(protagonistId, affectedId)
	if err != nil {
		return err
	}
	for _, event := range attackEvents {
		err := s.repo.Delete(event.AttackEventId)
		if err != nil {
			return err
		}
	}
	return nil
}


func (s *service) CreateEvent(attackEvent domain.AttackEvent) (domain.AttackEvent, error) {

	// eventDomain := domain.AttackEvent{
	// 	Type:               eventDto.Type,
	// 	Weapon:             eventDto.Weapon,
	// 	Spell:              eventDto.Spell,
	// 	Environment:        eventDto.Environment,
	// 	Session_id:         eventDto.Session_id,
	// 	EventProtagonistId: eventDto.EventProtagonistId,
	// 	EventResolution:    eventDto.EventResolution,
	// 	DmgType:            eventDto.DmgType,
	// 	Description:        eventDto.Description,
	// 	TimeStamp:          eventDto.TimeStamp,
	// }

	createdEvent, err := s.repo.Create(attackEvent)
	if err != nil {
		return domain.AttackEvent{}, err
	}

	return createdEvent, nil
}

func (s *service) GetAllEvents() ([]domain.AttackEvent, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *service) GetEventById(id int) (dto.ResponseEventDto, error) {
	event, err := s.repo.GetById(id)
	if err != nil {
		return dto.ResponseEventDto{}, err
	}

	session := domain.Session{
		SessionId:   event.Session_id,
		Start:       event.Start,
		End:         event.End,
		Description: event.SessionDescription,
		CampaignId:  &event.SessionCampaignId,
	}

	eventProtagonist := dto.CharacterCardDto{
		CharacterId: event.CharacterId,
		UserId:      event.CharacterUserId,
		CampaignID:  event.CharacterCampaignId,
		Name:        event.CharacterName,
		Race:        event.RaceName,
		Class:       event.ClassName,
		Level:       event.Level,
		HitPoints:   event.HitPoints,
	}

	affected, err := s.repo.GetCharacterDataByAttackEventId(event.AttackEventId)
	if err != nil {
		return dto.ResponseEventDto{}, err
	}

	eventToReturn := dto.ResponseEventDto{
		AttackEventId:    event.AttackEventId,
		Type:             event.Type,
		Environment:      event.Environment,
		Session:          session,
		EventProtagonist: eventProtagonist,
		EventResolution:  event.EventResolution,
		Weapon:           event.Weapon,
		Spell:            event.Spell,
		DmgType:          event.DmgType,
		Description:      event.Description,
		TimeStamp:        event.TimeStamp,
		Affected:         affected,
	}

	return eventToReturn, nil
}

func (s *service) GetEventsBySessionId(sessionid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetBySessionId(sessionid)
	if err != nil {
		return nil, err
	}

	results := make(chan dto.ResponseEventDto, len(events))
	errors := make(chan error, len(events))

	for _, event := range events {
		go func(event dto.RepositoryResponseAttackEvent) {
			session := domain.Session{
				SessionId:   event.Session_id,
				Start:       event.Start,
				End:         event.End,
				Description: event.SessionDescription,
				CampaignId:  &event.SessionCampaignId,
			}

			eventProtagonist := dto.CharacterCardDto{
				CharacterId: event.CharacterId,
				UserId:      event.CharacterUserId,
				CampaignID:  event.CharacterCampaignId,
				Name:        event.CharacterName,
				Race:        event.RaceName,
				Class:       event.ClassName,
				Level:       event.Level,
				HitPoints:   event.HitPoints,
			}

			affected, err := s.repo.GetCharacterDataByAttackEventId(event.AttackEventId)
			if err != nil {
				errors <- err
				return
			}

			eventToReturn := dto.ResponseEventDto{
				AttackEventId:    event.AttackEventId,
				Type:             event.Type,
				Environment:      event.Environment,
				Session:          session,
				EventProtagonist: eventProtagonist,
				EventResolution:  event.EventResolution,
				Weapon:           event.Weapon,
				Spell:            event.Spell,
				DmgType:          event.DmgType,
				Description:      event.Description,
				TimeStamp:        event.TimeStamp,
				Affected:         affected,
			}
			results <- eventToReturn
		}(event)
	}

	var eventsToReturn []dto.ResponseEventDto
	for i := 0; i < len(events); i++ {
		select {
		case eventToReturn := <-results:
			eventsToReturn = append(eventsToReturn, eventToReturn)
		case err := <-errors:
			return nil, err
		}
	}

	return eventsToReturn, nil
}

func (s *service) GetEventsByProtagonistId(protagonistid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetByProtagonistId(protagonistid)
	if err != nil {
		return nil, err
	}

	results := make(chan dto.ResponseEventDto, len(events))
	errors := make(chan error, len(events))

	for _, event := range events {
		go func(event dto.RepositoryResponseAttackEvent) {
			session := domain.Session{
				SessionId:   event.Session_id,
				Start:       event.Start,
				End:         event.End,
				Description: event.SessionDescription,
				CampaignId:  &event.SessionCampaignId,
			}

			eventProtagonist := dto.CharacterCardDto{
				CharacterId: event.CharacterId,
				UserId:      event.CharacterUserId,
				CampaignID:  event.CharacterCampaignId,
				Name:        event.CharacterName,
				Race:        event.RaceName,
				Class:       event.ClassName,
				Level:       event.Level,
				HitPoints:   event.HitPoints,
			}

			affected, err := s.repo.GetCharacterDataByAttackEventId(event.AttackEventId)
			if err != nil {
				errors <- err
				return
			}

			eventToReturn := dto.ResponseEventDto{
				AttackEventId:    event.AttackEventId,
				Type:             event.Type,
				Environment:      event.Environment,
				Session:          session,
				EventProtagonist: eventProtagonist,
				EventResolution:  event.EventResolution,
				Weapon:           event.Weapon,
				Spell:            event.Spell,
				DmgType:          event.DmgType,
				Description:      event.Description,
				TimeStamp:        event.TimeStamp,
				Affected:         affected,
			}
			results <- eventToReturn
		}(event)
	}

	var eventsToReturn []dto.ResponseEventDto
	for i := 0; i < len(events); i++ {
		select {
		case eventToReturn := <-results:
			eventsToReturn = append(eventsToReturn, eventToReturn)
		case err := <-errors:
			return nil, err
		}
	}

	return eventsToReturn, nil
}

func (s *service) GetEventsByAffectedId(affectedid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetByAffectedId(affectedid)
	if err != nil {
		return nil, err
	}

	results := make(chan dto.ResponseEventDto, len(events))
	errors := make(chan error, len(events))

	for _, event := range events {
		go func(event dto.RepositoryResponseAttackEvent) {
			session := domain.Session{
				SessionId:   event.Session_id,
				Start:       event.Start,
				End:         event.End,
				Description: event.SessionDescription,
				CampaignId:  &event.SessionCampaignId,
			}

			eventProtagonist := dto.CharacterCardDto{
				CharacterId: event.CharacterId,
				UserId:      event.CharacterUserId,
				CampaignID:  event.CharacterCampaignId,
				Name:        event.CharacterName,
				Race:        event.RaceName,
				Class:       event.ClassName,
				Level:       event.Level,
				HitPoints:   event.HitPoints,
			}

			affected, err := s.repo.GetCharacterDataByAttackEventId(event.AttackEventId)
			if err != nil {
				errors <- err
				return
			}

			eventToReturn := dto.ResponseEventDto{
				AttackEventId:    event.AttackEventId,
				Type:             event.Type,
				Environment:      event.Environment,
				Session:          session,
				EventProtagonist: eventProtagonist,
				EventResolution:  event.EventResolution,
				Weapon:           event.Weapon,
				Spell:            event.Spell,
				DmgType:          event.DmgType,
				Description:      event.Description,
				TimeStamp:        event.TimeStamp,
				Affected:         affected,
			}
			results <- eventToReturn
		}(event)
	}

	var eventsToReturn []dto.ResponseEventDto
	for i := 0; i < len(events); i++ {
		select {
		case eventToReturn := <-results:
			eventsToReturn = append(eventsToReturn, eventToReturn)
		case err := <-errors:
			return nil, err
		}
	}

	return eventsToReturn, nil
}

func (s *service) GetEventsByProtagonistIdAndAffectedId(protagonistid, affectedid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetByProtagonistIdAndAffectedId(protagonistid, affectedid)
	if err != nil {
		return nil, err
	}

	results := make(chan dto.ResponseEventDto, len(events))
	errors := make(chan error, len(events))

	for _, event := range events {
		go func(event dto.RepositoryResponseAttackEvent) {
			session := domain.Session{
				SessionId:   event.Session_id,
				Start:       event.Start,
				End:         event.End,
				Description: event.SessionDescription,
				CampaignId:  &event.SessionCampaignId,
			}

			eventProtagonist := dto.CharacterCardDto{
				CharacterId: event.CharacterId,
				UserId:      event.CharacterUserId,
				CampaignID:  event.CharacterCampaignId,
				Name:        event.CharacterName,
				Race:        event.RaceName,
				Class:       event.ClassName,
				Level:       event.Level,
				HitPoints:   event.HitPoints,
			}

			affected, err := s.repo.GetCharacterDataByAttackEventId(event.AttackEventId)
			if err != nil {
				errors <- err
				return
			}

			eventToReturn := dto.ResponseEventDto{
				AttackEventId:    event.AttackEventId,
				Type:             event.Type,
				Environment:      event.Environment,
				Session:          session,
				EventProtagonist: eventProtagonist,
				EventResolution:  event.EventResolution,
				Weapon:           event.Weapon,
				Spell:            event.Spell,
				DmgType:          event.DmgType,
				Description:      event.Description,
				TimeStamp:        event.TimeStamp,
				Affected:         affected,
			}
			results <- eventToReturn
		}(event)
	}

	var eventsToReturn []dto.ResponseEventDto
	for i := 0; i < len(events); i++ {
		select {
		case eventToReturn := <-results:
			eventsToReturn = append(eventsToReturn, eventToReturn)
		case err := <-errors:
			return nil, err
		}
	}

	return eventsToReturn, nil
}

func (s *service) UpdateEvent(eventDto dto.CreateAttackEventDto, id int) (domain.AttackEvent, error) {

	timestamp := time.Now()

	eventDomain := domain.AttackEvent{
		Type:               eventDto.Type,
		Weapon:             eventDto.Weapon,
		Spell:              eventDto.Spell,
		Environment:        eventDto.Environment,
		Session_id:         eventDto.Session_id,
		EventProtagonistId: eventDto.EventProtagonistId,
		EventResolution:    eventDto.EventResolution,
		DmgType:            eventDto.DmgType,
		Description:        eventDto.Description,
		TimeStamp:          &timestamp,
	}

	updatedEvent, err := s.repo.Update(eventDomain, id)
	if err != nil {
		return domain.AttackEvent{}, err
	}

	return updatedEvent, nil
}

func (s *service) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
