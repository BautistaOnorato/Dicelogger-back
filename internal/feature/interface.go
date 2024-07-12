package feature

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type FeatureHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetAllFeaturesByCharacterId() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type FeatureService interface {
	CreateFeature (dto.CreateFeatureDto) (domain.Feature, error)
	GetAllFeatures() ([]domain.Feature, error)
	GetAllFeaturesByCharacterId(characterId int) (dto.FeatureFullResponseDto, error)
	GetFeatureById(id int) (domain.Feature, error)
	UpdateFeature(feature dto.CreateFeatureDto, id int) (domain.Feature, error)
	DeleteFeature(id int) error
}

type FeatureRepository interface {
	Create(feature dto.CreateFeatureDto) (domain.Feature, error)
	GetAll() ([]domain.Feature, error)
	GetAllByCharacterId(characterId int) ([]domain.Feature, error)
	GetById(id int) (domain.Feature, error)
	Update(feature domain.Feature, id int) (domain.Feature, error)
	Delete(id int) error
}