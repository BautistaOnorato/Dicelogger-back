package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	characterXproficiency "github.com/proyecto-dnd/backend/internal/characterXProficiency"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type CharacterXProficiencyHandler struct {
	service characterXproficiency.CharacterXProficiencyService
}

func NewCharacterXProficiencyHandler(service characterXproficiency.CharacterXProficiencyService) *CharacterXProficiencyHandler {
	return &CharacterXProficiencyHandler{service: service}
}

func (h *CharacterXProficiencyHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCharacterXProficiency domain.CharacterXProficiency
		if err := ctx.BindJSON(&tempCharacterXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdCharacterXProficiency, err := h.service.Create(tempCharacterXProficiency)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdCharacterXProficiency)
	}
}

func (h *CharacterXProficiencyHandler) HandlerDelete() gin.HandlerFunc {
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
		ctx.JSON(200, "Deleted characterXproficiency with id "+id)
	}
}
