package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/attackEvent"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type AttackEventHandler struct {
	service attackEvent.AttackEventService
}

func NewAttackEventHandler(service *attackEvent.AttackEventService) *AttackEventHandler {
	return &AttackEventHandler{service: *service}
}

// event godoc
// @Summary Create event
// @Tags event
// @Accept json
// @Produce json
// @Param body body dto.CreateAttackEventDto true "CreateEventDto"
// @Success 201 {object} domain.Event
// @Failure 500 {object} error
// @Router /event [post]
func (h *AttackEventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempEvent domain.AttackEvent
		if err := ctx.BindJSON(&tempEvent); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		createdEvent, err := h.service.CreateEvent(tempEvent)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdEvent)
	}
}

// event godoc
// @Summary Get all events
// @Tags event
// @Produce json
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event [get]
func (h *AttackEventHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventList, err := h.service.GetAllEvents()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Get event by id
// @Tags event
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/{id} [get]
func (h *AttackEventHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempEvent, err := h.service.GetEventById(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempEvent)
	}
}

// event godoc
// @Summary Get event by session id
// @Tags event
// @Produce json
// @Param id path int true "session_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/session/{id} [get]
func (h *AttackEventHandler) HandlerGetBySessionId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		eventList, err := h.service.GetEventsBySessionId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Get event by protagonist id
// @Tags event
// @Produce json
// @Param id path int true "protagonist_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/protagonist/{id} [get]
func (h *AttackEventHandler) HandlerGetByProtagonistId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		eventList, err := h.service.GetEventsByProtagonistId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Get event by affected id
// @Tags event
// @Produce json
// @Param id path int true "affected_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/affected/{id} [get]
func (h *AttackEventHandler) HandlerGetByAffectedId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		eventList, err := h.service.GetEventsByAffectedId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Get event by protagonist id and affected id
// @Tags event
// @Produce json
// @Param id path int true "protagonist_id"
// @Param id path int true "affected_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/protagonist/{protagonistid}/affected/{affectedid} [get]
func (h *AttackEventHandler) HandlerGetByProtagonistIdAndAffectedId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		protagonistid := ctx.Param("protagonistid")
		affectedid := ctx.Param("affectedid")

		intProtagonistId, err := strconv.Atoi(protagonistid)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		intAffectedId, err := strconv.Atoi(affectedid)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		eventList, err := h.service.GetEventsByProtagonistIdAndAffectedId(intProtagonistId, intAffectedId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Update event
// @Tags event
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.CreateAttackEventDto true "CreateEventDto"
// @Success 200 {object} domain.Event
// @Failure 500 {object} error
// @Router /event/{id} [put]
func (h *AttackEventHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var tempEvent dto.CreateAttackEventDto
		if err := ctx.BindJSON(&tempEvent); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		updatedEvent, err := h.service.UpdateEvent(tempEvent, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, updatedEvent)
	}
}

// event godoc
// @Summary Delete event
// @Tags event
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /event/{id} [delete]
func (h *AttackEventHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		serviceErr := h.service.DeleteEvent(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr.Error())
			return
		}

		ctx.JSON(200, "Deleted Event with id "+id)
	}
}
