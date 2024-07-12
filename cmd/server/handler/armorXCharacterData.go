package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/armorXCharacterData"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type ArmorXCharacterDataHandler struct {
	service armorXCharacterData.ServiceArmorXCharacterData
}

func NewArmorXCharacterDataHandler(service *armorXCharacterData.ServiceArmorXCharacterData) *ArmorXCharacterDataHandler {
	return &ArmorXCharacterDataHandler{service: *service}
}

// armorXCharacterData godoc
// @Summary Create armorXCharacterData
// @Tags armorXCharacterData
// @Accept json
// @Produce json
// @Param body body domain.ArmorXCharacterData true "ArmorXCharacterData"
// @Success 201 {object} domain.ArmorXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /armor_character [post]
func (h *ArmorXCharacterDataHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempArmorXCharacterData domain.ArmorXCharacterData
		if err := ctx.BindJSON(&tempArmorXCharacterData); err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		createdArmorXCharacterData, err := h.service.CreateArmorXCharacterData(tempArmorXCharacterData)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(201, createdArmorXCharacterData)
	}
}

// armorXCharacterData godoc
// @Summary Delete armorXCharacterData
// @Tags armorXCharacterData
// @Produce json
// @Param id path int true "id"
// @Success 204 {object} nil
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /armor_character/{id} [delete]
func (h *ArmorXCharacterDataHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		err = h.service.DeleteArmorXCharacterData(id)
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(204, nil)
	}
}

// armorXCharacterData godoc
// @Summary Delete armorXCharacterData by character id
// @Tags armorXCharacterData
// @Produce json
// @Param id path int true "character_id"
// @Success 204 {object} nil
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /armor_character/character/{id} [delete]
func (h *ArmorXCharacterDataHandler) HandlerDeleteByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		err = h.service.DeleteByCharacterDataIdArmor(id)
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(204, nil)
	}
}

// armorXCharacterData godoc
// @Summary Get all itemXCharacterData
// @Tags armorXCharacterData
// @Produce json
// @Success 200 {array} domain.ArmorXCharacterData
// @Failure 500 {object} error
// @Router /armor_character [get]
func (h *ArmorXCharacterDataHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		armorXCharacterDataList, err := h.service.GetAllArmorXCharacterData()
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, armorXCharacterDataList)
	}
}

// armorXCharacterData godoc
// @Summary Get armorXCharacterData by character id
// @Tags armorXCharacterData
// @Produce json
// @Param id path int true "character_id"
// @Success 200 {array} domain.ArmorXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /armor_character/character/{id} [get]
func (h *ArmorXCharacterDataHandler) HandlerGetByCharacterDataId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		armorXCharacterDataList, err := h.service.GetByCharacterDataIdArmor(id)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, armorXCharacterDataList)
	}
}

// armorXCharacterData godoc
// @Summary Get armorXCharacterData by id
// @Tags armorXCharacterData
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.ArmorXCharacterData
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /armor_character/{id} [get]
func (h *ArmorXCharacterDataHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		armorXCharacterData, err := h.service.GetByIdArmorXCharacterData(id)
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(200, armorXCharacterData)
	}
}

// armorXCharacterData godoc
// @Summary Update armorXCharacterData
// @Tags armorXCharacterData
// @Accept json
// @Produce json
// @Param body body domain.ArmorXCharacterData true "ArmorXCharacterData"
// @Param id path int true "id"
// @Success 200 {object} domain.ArmorXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /armor_character/{id} [put]
func (h *ArmorXCharacterDataHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempArmorXCharacterData domain.ArmorXCharacterData
		if err := ctx.BindJSON(&tempArmorXCharacterData); err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}

		tempArmorXCharacterData.ArmorXCharacterData_Id = id

		updatedItemXCharacterData, err := h.service.UpdateArmorXCharacterData(tempArmorXCharacterData)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, updatedItemXCharacterData)
	}
}