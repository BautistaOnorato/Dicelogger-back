package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/spell"
)

type SpellHandler struct {
	service spell.ServiceSpell
}

func NewSpellHandler(service *spell.ServiceSpell) *SpellHandler {
	return &SpellHandler{service: *service}
}

func (h *SpellHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempSpell dto.SpellDto
		if err := ctx.BindJSON(&tempSpell); err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		createdSpell, err := h.service.Create(tempSpell)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdSpell)
	}
}

func (h *SpellHandler) HandlergetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spellList, err := h.service.GetAll()
		if err != nil {

			ctx.JSON(400, err.Error())
			return
		}

		ctx.JSON(200, spellList)
	}
}

func (h *SpellHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempSpell, err := h.service.GetById(intId)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		ctx.JSON(200, tempSpell)
	}
}

func (h *SpellHandler) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		spells, err := h.service.GetByCharacterDataId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, spells)
	}
}

func (h *SpellHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var tempSpell dto.SpellDto
		if err := ctx.BindJSON(&tempSpell); err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		updatedSpell, err := h.service.Update(tempSpell, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, updatedSpell)
	}
}

func (h *SpellHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		if err = h.service.Delete(intId); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, "Deleted spell with id: "+id)
	}
}

func (h *SpellHandler) HandlerGetByClassId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		spells, err := h.service.GetByClassId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, spells)
	}
}
