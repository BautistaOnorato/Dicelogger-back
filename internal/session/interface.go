package session

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type SessionHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetByCampaignId() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type SessionService interface {
	CreateSession(dto.CreateSessionDto) (domain.Session, error)
	GetAllSessions() ([]domain.Session, error)
	GetSessionById(id int) (domain.Session, error)
	GetSessionsByCampaignId(id int) ([]domain.Session, error)
	UpdateSession(Session dto.CreateSessionDto, id int) (domain.Session, error)
	DeleteSession(id int) error
}

type SessionRepository interface {
	Create(Session domain.Session) (domain.Session, error)
	GetAll()([]domain.Session, error)
	GetById(id int)(domain.Session, error)
	GetByCampaignId(id int)([]domain.Session, error)
	Update(Session domain.Session, id int) (domain.Session, error)
	Delete(id int)error
}