package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/background"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type BackgroundHandler struct {
	service background.BackgroundService
}

func NewBackgroundHandler(service background.BackgroundService) *BackgroundHandler {
	return &BackgroundHandler{service: service}
}

func (h *BackgroundHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempBackground dto.CreateBackgroundDto
		if err := ctx.BindJSON(&tempBackground); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		createdBackground, err := h.service.CreateBackground(tempBackground)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdBackground)
	}
}

func (h *BackgroundHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		backgroundList, err := h.service.GetAllBackgrounds()
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, backgroundList)
	}
}

func (h *BackgroundHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempBackground, err := h.service.GetBackgroundByID(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempBackground)
	}
}

func (h *BackgroundHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		var tempBackground dto.CreateBackgroundDto
		if err := ctx.BindJSON(&tempBackground); err != nil {

			ctx.JSON(500, err.Error())
			return
		}

		updatedBackground, err := h.service.UpdateBackground(tempBackground, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, updatedBackground)
	}
}

func (h *BackgroundHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		serviceErr := h.service.DeleteBackground(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr.Error())
			return
		}

		ctx.JSON(200, "Deleted Background with id "+id)
	}
}
