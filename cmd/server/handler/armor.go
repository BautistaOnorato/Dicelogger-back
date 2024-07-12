package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/armor"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type ArmorHandler struct {
	service armor.ArmorService
}

func NewArmorHandler(service *armor.ArmorService) *ArmorHandler {
	return &ArmorHandler{service: *service}
}

func (h *ArmorHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempArmor dto.CreateArmorDto
		if err := ctx.BindJSON(&tempArmor); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		createdArmor, err := h.service.CreateArmor(tempArmor)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdArmor)
	}
}

func (h *ArmorHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		armorList, err := h.service.GetAllArmor()
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, armorList)
	}
}

func (h *ArmorHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempArmor, err := h.service.GetArmorByID(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempArmor)
	}
}

func (h *ArmorHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var tempArmor dto.CreateArmorDto
		if err := ctx.BindJSON(&tempArmor); err != nil {

			ctx.JSON(500, err.Error())
			return
		}

		updatedArmor, err := h.service.UpdateArmor(tempArmor, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, updatedArmor)
	}
}

func (h *ArmorHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		serviceErr := h.service.DeleteArmor(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr.Error())
			return
		}

		ctx.JSON(200, "Deleted Armor with id "+id)
	}
}
