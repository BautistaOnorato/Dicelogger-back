package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/race"
)

type RaceHandler struct {
	service race.RaceService
}

func NewRaceHandler(service race.RaceService) *RaceHandler {
	return &RaceHandler{service: service}
}

func (h *RaceHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempRace dto.CreateRaceDto
		if err := ctx.BindJSON(&tempRace); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdRace, err := h.service.CreateRace(tempRace)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdRace)
	}
}

func (h *RaceHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		races, err := h.service.GetAllRaces()
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, races)
	}
}

func (h *RaceHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		race, err := h.service.GetRaceByID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, race)
	}
}

func (h *RaceHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		var tempRace dto.CreateRaceDto
		if err := ctx.BindJSON(&tempRace); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		updatedRace, err := h.service.UpdateRace(tempRace, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedRace)
	}
}

func (h *RaceHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		serviceErr := h.service.DeleteRace(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr)
			return
		}

		ctx.JSON(200, "Deleted Race with id "+id)
	}
}
