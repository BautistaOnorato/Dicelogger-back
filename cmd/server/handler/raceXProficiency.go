package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	raceXproficiency "github.com/proyecto-dnd/backend/internal/raceXProficiency"
)

type RaceXProficiencyHandler struct {
	service raceXproficiency.RaceXProficiencyService
}

func NewRaceXProficiencyHandler(service *raceXproficiency.RaceXProficiencyService) *RaceXProficiencyHandler {
	return &RaceXProficiencyHandler{service: *service}
}

func (h *RaceXProficiencyHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempRaceXProficiency domain.RaceXProficiency
		if err := ctx.BindJSON(&tempRaceXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdRaceXProficiency, err := h.service.Create(tempRaceXProficiency)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdRaceXProficiency)
	}
}

func (h *RaceXProficiencyHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempRaceXProficiency domain.RaceXProficiency
		if err := ctx.BindJSON(&tempRaceXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		if err := h.service.Delete(tempRaceXProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Successfully deleted row with race_id="+strconv.Itoa(tempRaceXProficiency.RaceId)+" and proficiency_id="+strconv.Itoa(tempRaceXProficiency.ProficiencyId))
	}
}
