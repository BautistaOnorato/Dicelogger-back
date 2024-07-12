package attackEvent

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type AttackEventHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetBySessionId() gin.HandlerFunc
	HandlerGetByProtagonistId() gin.HandlerFunc
	HandlerGetByAffectedId() gin.HandlerFunc
	HandlerGetByProtagonistIdAndAffectedId() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type AttackEventService interface {
	CreateEvent(domain.AttackEvent) (domain.AttackEvent, error)
	GetAllEvents() ([]domain.AttackEvent, error)
	GetEventById(id int) (dto.ResponseEventDto, error)
	GetEventsBySessionId(sessionid int) ([]dto.ResponseEventDto, error)
	GetEventsByProtagonistId(protagonistId int) ([]dto.ResponseEventDto, error)
	GetEventsByAffectedId(affectedId int) ([]dto.ResponseEventDto, error)
	GetEventsByProtagonistIdAndAffectedId(protagonistId int, affectedId int) ([]dto.ResponseEventDto, error)
	UpdateEvent(event dto.CreateAttackEventDto, id int) (domain.AttackEvent, error)
	DeleteEvent(id int) error
	DeleteByProtagonistAndAffectedId(protagonistId int, affectedId int) error
}

type AttackEventRepository interface {
	Create(event domain.AttackEvent) (domain.AttackEvent, error)
	GetAll() ([]domain.AttackEvent, error)
	GetById(id int) (dto.RepositoryResponseAttackEvent, error)
	GetBySessionId(sessionid int) ([]dto.RepositoryResponseAttackEvent, error)
	GetByProtagonistId(protagonistid int) ([]dto.RepositoryResponseAttackEvent, error)
	GetByAffectedId(affectedid int) ([]dto.RepositoryResponseAttackEvent, error)
	GetCharacterDataByAttackEventId (id int) ([]dto.CharacterCardDto, error)
	GetByProtagonistIdAndAffectedId(protagonistid int, affectedid int) ([]dto.RepositoryResponseAttackEvent, error)
	Update(event domain.AttackEvent, id int) (domain.AttackEvent, error)
	Delete(id int) error
}
