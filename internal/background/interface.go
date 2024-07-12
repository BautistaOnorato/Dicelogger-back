package background

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type BackgroundHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type BackgroundService interface {
	CreateBackground(dto.CreateBackgroundDto) (domain.Background, error)
	GetAllBackgrounds() ([]domain.Background, error)
	GetBackgroundByID(id int) (domain.Background, error)
	UpdateBackground(background dto.CreateBackgroundDto, id int) (domain.Background, error)
	DeleteBackground(id int) error
}

type BackgroundRepository interface {
	Create(domain.Background) (domain.Background, error)
	GetAllBackgrounds() ([]domain.Background, error)
	GetBackgroundById(id int) (domain.Background, error)
	UpdateBackground(background domain.Background, id int) (domain.Background, error)
	DeleteBackground(id int) error
}
