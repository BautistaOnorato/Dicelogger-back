package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/character_feature"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterFeature struct {
	service character_feature.CharacterFeatureService
}

func NewCharacterFeatureHandler(service *character_feature.CharacterFeatureService) *CharacterFeature {
	return &CharacterFeature{service: *service}
}

// characterFeature godoc
// @Summary Create characterFeature
// @Tags characterFeature
// @Accept json
// @Produce json
// @Param body body dto.CreateCharacterFeatureDto true "CreateCharacterFeatureDto"
// @Success 201 {object} domain.CharacterFeature
// @Failure 500 {object} error
// @Router /character_feature [post]
func (h *CharacterFeature) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCharacterFeature dto.CreateCharacterFeatureDto
		if err := ctx.BindJSON(&tempCharacterFeature); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdCharacterFeature, err := h.service.CreateCharacterFeature(tempCharacterFeature)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdCharacterFeature)
	}
}

// characterFeature godoc
// @Summary Get all characterFeatures
// @Tags characterFeature
// @Produce json
// @Success 200 {array} domain.CharacterFeature
// @Failure 500 {object} error
// @Router /character_feature [get]
func (h *CharacterFeature) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterFeatures, err := h.service.GetAllCharacterFeatures()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, characterFeatures)
	}
}

// characterFeature godoc
// @Summary Get characterFeature by feature id
// @Tags characterFeature
// @Produce json
// @Param id path int true "feature_id"
// @Success 200 {array} domain.CharacterFeature
// @Failure 500 {object} error
// @Router /character_feature/feature/{id} [get]
func (h *CharacterFeature) HandlerGetByFeatureId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempCharacterFeature, err := h.service.GetCharacterFeatureByFeatureId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempCharacterFeature)
	}
}

// characterFeature godoc
// @Summary Get characterFeature by character id
// @Tags characterFeature
// @Produce json
// @Param id path int true "character_id"
// @Success 200 {array} domain.CharacterFeature
// @Failure 500 {object} error
// @Router /character_feature/character/{id} [get]
func (h *CharacterFeature) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempCharacterFeature, err := h.service.GetCharacterFeatureByCharacterId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempCharacterFeature)
	}
}

// characterFeature godoc
// @Summary Delete characterFeature
// @Tags characterFeature
// @Produce json
// @Param body body dto.CreateCharacterFeatureDto true "CreateCharacterFeatureDto"
// @Success 204
// @Failure 500 {object} error
// @Router /character_feature [delete]
func (h *CharacterFeature) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idFeature, err := strconv.Atoi(ctx.Param("idFeature"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		idCharacter, err := strconv.Atoi(ctx.Param("idCharacter"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		err = h.service.DeleteCharacterFeature(idFeature, idCharacter)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(204, nil)
	}
}

func (h *CharacterFeature) HandlerDeleteQueries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idCharacter, err := strconv.Atoi(ctx.Query("idCharacter"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		idFeature, err := strconv.Atoi(ctx.Query("idFeature"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		err = h.service.DeleteCharacterFeature(idFeature, idCharacter)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "The row with idCharacter = "+strconv.Itoa(idCharacter)+" and idFeature = "+strconv.Itoa(idFeature)+" was deleted")
	}
}
