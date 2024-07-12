package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	characterXAttackEvent "github.com/proyecto-dnd/backend/internal/characterXAttackEvent"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterXAttackEventHandler struct {
	service characterXAttackEvent.CharacterXAttackEventService
}

func NewCharacterXAttackEventHandler(service characterXAttackEvent.CharacterXAttackEventService) *CharacterXAttackEventHandler {
	return &CharacterXAttackEventHandler{service: service}
}

func (h *CharacterXAttackEventHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterXSpellEvents, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEvents)
	}
}

func (h *CharacterXAttackEventHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		characterXSpellEventId, err := h.service.GetById(idInt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEventId)
	}
}

func (h *CharacterXAttackEventHandler) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterId := ctx.Param("characterId")

		characterIdInt, err := strconv.Atoi(characterId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		characterXSpellEvents, err := h.service.GetByCharacterId(characterIdInt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEvents)
	}
}

func (h *CharacterXAttackEventHandler) HandlerGetByEventId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spellEventId := ctx.Param("spellEventId")

		spellEventIdInt, err := strconv.Atoi(spellEventId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		characterXSpellEvents, err := h.service.GetByEventId(spellEventIdInt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEvents)
	}
}

func (h *CharacterXAttackEventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCharacterXSpellEvent dto.CharacterXAttackEventDto
		if err := ctx.BindJSON(&tempCharacterXSpellEvent); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdCharacterXSpellEvent, err := h.service.Create(tempCharacterXSpellEvent)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdCharacterXSpellEvent)
	}
}

func (h *CharacterXAttackEventHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		err = h.service.Delete(int(intId))
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Deleted characterXattackevent with id "+id)
	}
}