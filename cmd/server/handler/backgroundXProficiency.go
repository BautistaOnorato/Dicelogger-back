package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	backgroundXproficiency "github.com/proyecto-dnd/backend/internal/backgroundXProficiency"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type BackgroundXProficiencyHandler struct {
	service backgroundXproficiency.BackgroundXProficiencyService
}

func NewBackgroundXProficiencyHandler(service backgroundXproficiency.BackgroundXProficiencyService) *BackgroundXProficiencyHandler {
	return &BackgroundXProficiencyHandler{service: service}
}

func (h *BackgroundXProficiencyHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempBackgroundXProficiency domain.BackgroundXProficiency
		if err := ctx.BindJSON(&tempBackgroundXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdBackgroundXProficiency, err := h.service.Create(tempBackgroundXProficiency)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdBackgroundXProficiency)
	}
}

func (h *BackgroundXProficiencyHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempBackgroundXProficiency domain.BackgroundXProficiency
		if err := ctx.BindJSON(&tempBackgroundXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		if err := h.service.Delete(tempBackgroundXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Successfully deleted row with background_id="+strconv.Itoa(tempBackgroundXProficiency.BackgroundID)+" and proficiency_id="+strconv.Itoa(tempBackgroundXProficiency.ProficiencyID))
	}
}
