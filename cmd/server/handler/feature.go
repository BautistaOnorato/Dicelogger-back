package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/feature"
)

type FeatureHandler struct {
	service feature.FeatureService
}

func NewFeatureHandler(service *feature.FeatureService) *FeatureHandler {
	return &FeatureHandler{service: *service}
}

// feature godoc
// @Summary Create feature
// @Tags feature
// @Accept json
// @Produce json
// @Param body body dto.CreateFeatureDto true "CreateFeatureDto"
// @Success 201 {object} domain.Feature
// @Failure 500 {object} error
// @Router /feature [post]
func (h *FeatureHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempFeature dto.CreateFeatureDto
		if err := ctx.BindJSON(&tempFeature); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdFeature, err := h.service.CreateFeature(tempFeature)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdFeature)
	}
}

// feature godoc
// @Summary Get all features
// @Tags feature
// @Produce json
// @Success 200 {array} domain.Feature
// @Failure 500 {object} error
// @Router /feature [get]
func (h *FeatureHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		featureList, err := h.service.GetAllFeatures()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, featureList)
	}
}

// feature godoc
// @Summary Get all features by character id
// @Tags feature
// @Produce json
// @Param id path int true "id"
// @Success 200 {array} dto.FeatureFullResponseDto
// @Failure 500 {object} error
// @Router /feature/character/{id} [get]
func (h *FeatureHandler) HandlerGetAllFeaturesByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		featureList, err := h.service.GetAllFeaturesByCharacterId(characterId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, featureList)
	}
}

// feature godoc
// @Summary Get feature by id
// @Tags feature
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Feature
// @Failure 500 {object} error
// @Router /feature/{id} [get]
func (h *FeatureHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempFeature, err := h.service.GetFeatureById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempFeature)
	}
}

// feature godoc
// @Summary Update feature
// @Tags feature
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.CreateFeatureDto true "CreateFeatureDto"
// @Success 201 {object} domain.Feature
// @Failure 500 {object} error
// @Router /feature/{id} [put]
func (h *FeatureHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		var tempFeature dto.CreateFeatureDto
		if err := ctx.BindJSON(&tempFeature); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdFeature, err := h.service.UpdateFeature(tempFeature, id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdFeature)
	}
}

// feature godoc
// @Summary Delete feature
// @Tags feature
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /feature/{id} [delete]
func (h *FeatureHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		err = h.service.DeleteFeature(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "Feature deleted")
	}
}

func(h *FeatureHandler) HandlerDeleteByQueries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		err = h.service.DeleteFeature(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "Feature deleted")
	}
}