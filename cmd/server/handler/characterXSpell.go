package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	characterXspell "github.com/proyecto-dnd/backend/internal/characterXSpell"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type CharacterXSpellHandler struct {
	service characterXspell.ServiceCharacterXSpell
}

func NewCharacterXSpellHandler(service *characterXspell.ServiceCharacterXSpell) *CharacterXSpellHandler {
	return &CharacterXSpellHandler{service: *service}
}

func (h *CharacterXSpellHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var tempCharacterXSpell domain.CharacterXSpell

		if err := ctx.BindJSON(&tempCharacterXSpell); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		createdCharacterXSpell, err := h.service.Create(tempCharacterXSpell)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdCharacterXSpell)
	}
}

func (h *CharacterXSpellHandler) HandlerDelete() gin.HandlerFunc {
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

		ctx.JSON(200, "Deleted characterXspell with id "+id)
	}
}

func (h *CharacterXSpellHandler) HandlerDeleteParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterId, err := strconv.Atoi(ctx.Query("characterId"))
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		spellId, err := strconv.Atoi(ctx.Query("spellId"))
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		err = h.service.DeleteParams(characterId, spellId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "Deleted characterXspell with characterId="+strconv.Itoa(characterId)+" and spellId="+strconv.Itoa(spellId))
	}
}
