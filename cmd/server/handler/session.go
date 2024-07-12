package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/session"
)

type SessionHandler struct {
	service session.SessionService
}

func NewSessionHandler(service *session.SessionService) *SessionHandler {
	return &SessionHandler{service: *service}
}

// session godoc
// @Summary Create session
// @Tags session
// @Accept json
// @Produce json
// @Param body body dto.CreateSessionDto true "CreateSessionDto"
// @Success 201 {object} domain.Session
// @Failure 500 {object} error
// @Router /session [post]
func (h *SessionHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempSession dto.CreateSessionDto
		if err := ctx.BindJSON(&tempSession); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdSession, err := h.service.CreateSession(tempSession)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdSession)
	}
}

// session godoc
// @Summary Get all sessions
// @Tags session
// @Produce json
// @Success 200 {array} domain.Session
// @Failure 500 {object} error
// @Router /session [get]
func (h *SessionHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionList, err := h.service.GetAllSessions()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, sessionList)
	}
}

// session godoc
// @Summary Get session by id
// @Tags session
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Session
// @Failure 500 {object} error
// @Router /session/{id} [get]
func (h *SessionHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempSession, err := h.service.GetSessionById(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempSession)
	}
}

// session godoc
// @Summary Get session by campaign id
// @Tags session
// @Produce json
// @Param id path int true "campaign_id"
// @Success 200 {array} domain.Session
// @Failure 500 {object} error
// @Router /session/campaign/{id} [get]
func (h *SessionHandler) HandlerGetByCampaignId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempSession, err := h.service.GetSessionsByCampaignId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempSession)
	}
}

// session godoc
// @Summary Update session
// @Tags session
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.CreateSessionDto true "CreateSessionDto"
// @Success 200 {object} domain.Session
// @Failure 500 {object} error
// @Router /session/{id} [put]
func (h *SessionHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var tempSession dto.CreateSessionDto
		if err := ctx.BindJSON(&tempSession); err != nil {

			ctx.JSON(500, err.Error())
			return
		}

		updatedSession, err := h.service.UpdateSession(tempSession, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, updatedSession)
	}
}

// session godoc
// @Summary Delete session
// @Tags session
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /session/{id} [delete]
func (h *SessionHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		serviceErr := h.service.DeleteSession(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr.Error())
			return
		}

		ctx.JSON(200, "Deleted Session with id "+id)
	}
}
