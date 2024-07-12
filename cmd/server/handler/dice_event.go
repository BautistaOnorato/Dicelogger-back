package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dice_event"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type DiceEventHandler struct {
	service dice_event.DiceEventService
}

func NewDiceEventHandler(service dice_event.DiceEventService) *DiceEventHandler {
	return &DiceEventHandler{service: service}
}

func (h *DiceEventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempDiceEvent domain.DiceEvent
		if err := ctx.BindJSON(&tempDiceEvent); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdDiceEvent, err := h.service.Create(tempDiceEvent)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdDiceEvent)
	}
}

func (h *DiceEventHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		diceEvents, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, diceEvents)
	}
}

func (h *DiceEventHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		diceEvent, err := h.service.GetById(id)
		if err != nil {
			ctx.JSON(404, err)
			return
		}
		ctx.JSON(200, diceEvent)
	}
}

func (h *DiceEventHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		var tempDiceEvent domain.DiceEvent
		if err := ctx.BindJSON(&tempDiceEvent); err != nil {
			ctx.JSON(500, err)
			return
		}
		tempDiceEvent.DiceEventId = id
		createdDiceEvent, err := h.service.Update(tempDiceEvent, id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdDiceEvent)
	}
}

func (h *DiceEventHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		err = h.service.Delete(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "deleted dice event with id: "+strconv.Itoa(id))
	}
}
