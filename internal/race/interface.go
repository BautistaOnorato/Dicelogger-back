package race

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RaceHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type RaceService interface {
	CreateRace(dto.CreateRaceDto) (domain.Race, error)
	GetAllRaces() ([]domain.Race, error)
	GetRaceByID(id int) (domain.Race, error)
	UpdateRace(race dto.CreateRaceDto, id int) (domain.Race, error)
	DeleteRace(id int) error
}

type RaceRepository interface {
	Create(race domain.Race) (domain.Race, error)
	GetAllRaces() ([]domain.Race, error)
	GetRaceById(id int) (domain.Race, error)
	UpdateRace(race domain.Race, id int) (domain.Race, error)
	DeleteRace(id int) error
}
