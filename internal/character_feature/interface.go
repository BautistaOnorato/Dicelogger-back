package character_feature

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterFeatureHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetByFeatureId() gin.HandlerFunc
	HandlerGetByCharacterId() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type CharacterFeatureService interface {
	CreateCharacterFeature(dto.CreateCharacterFeatureDto) (domain.CharacterFeature, error)
	GetAllCharacterFeatures() ([]domain.CharacterFeature, error)
	GetCharacterFeatureByFeatureId(id int) ([]domain.CharacterFeature, error)
	GetCharacterFeatureByCharacterId(id int) ([]domain.CharacterFeature, error)
	DeleteCharacterFeature(idFeature int, idCharacter int) error
	DeleteByCharacterDataId(idCharacter int) error
}

type CharacterFeatureRepository interface {
	Create(CharacterFeature domain.CharacterFeature) (domain.CharacterFeature, error)
	GetAll()([]domain.CharacterFeature, error)
	GetByFeatureId(id int)([]domain.CharacterFeature, error)
	GetByCharacterId(id int)([]domain.CharacterFeature, error)
	Delete(idFeature int, idCharacter int)error
	DeleteByCharacterDataId(idCharacter int)error
}