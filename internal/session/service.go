package session

import (
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	sessionRepo  SessionRepository
 }

func NewSessionService(sessionRepo SessionRepository) SessionService {
	return &service{sessionRepo: sessionRepo}
}

func (s *service) CreateSession(sessionDto dto.CreateSessionDto) (domain.Session, error) {
	sessionDomain := domain.Session{
		Start:       sessionDto.Start,
		End:         sessionDto.End,
		Description: sessionDto.Description,
		CampaignId:  sessionDto.CampaignId,
		CurrentEnviroment: sessionDto.CurrentEnviroment,
	}

	createdSession, err := s.sessionRepo.Create(sessionDomain)
	if err != nil {
		return domain.Session{}, err
	}

	return createdSession, nil
}

func (s *service) GetAllSessions() ([]domain.Session, error) {
	sessions, err := s.sessionRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *service) GetSessionById(id int) (domain.Session, error) {
	session, err := s.sessionRepo.GetById(id)
	if err != nil {
		return domain.Session{}, err
	}

	return session, nil
}

func (s *service) GetSessionsByCampaignId(id int) ([]domain.Session, error) {
	sessions, err := s.sessionRepo.GetByCampaignId(id)
	if err != nil {
		return nil, err
	}

	return sessions, nil

}

func (s *service) UpdateSession(sessionDto dto.CreateSessionDto, id int) (domain.Session, error) {

	sessionDomain := domain.Session{
		Start:       sessionDto.Start,
		End:         sessionDto.End,
		Description: sessionDto.Description,
		CampaignId:  sessionDto.CampaignId,
		CurrentEnviroment: sessionDto.CurrentEnviroment,
	}

	updatedSession, err := s.sessionRepo.Update(sessionDomain, id)
	if err != nil {
		fmt.Println(err)
		return domain.Session{}, err
	}

	return updatedSession, nil
}

func (s *service) DeleteSession(id int) error {
	return s.sessionRepo.Delete(id)
}
