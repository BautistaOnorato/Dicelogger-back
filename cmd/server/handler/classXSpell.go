package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	classXspell "github.com/proyecto-dnd/backend/internal/classXSpell"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type ClassXSpellHandler struct {
	service classXspell.ServiceClassXSpell
}

func NewClassXSpellHandler(service *classXspell.ServiceClassXSpell) *ClassXSpellHandler {
	return &ClassXSpellHandler{service: *service}
}

func (h *ClassXSpellHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempClassXSpell domain.ClassXSpell
		if err := ctx.BindJSON(&tempClassXSpell); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdClassXSpell, err := h.service.Create(tempClassXSpell)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdClassXSpell)
	}
}

func (h *ClassXSpellHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempClassXSpell domain.ClassXSpell
		if err := ctx.BindJSON(&tempClassXSpell); err != nil {
			ctx.JSON(500, err)
			return
		}
		if err := h.service.Delete(tempClassXSpell); err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, "Successfully deleted row with class_id="+strconv.Itoa(tempClassXSpell.ClassId)+" and spell_id="+strconv.Itoa(tempClassXSpell.SpellId))
	}
}
