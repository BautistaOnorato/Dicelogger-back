package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/class"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type ClassHandler struct {
	service class.ClassService
}

func NewClassHandler(service *class.ClassService) *ClassHandler {
	return &ClassHandler{service: *service}
}

// class godoc
// @Summary Create class
// @Tags class
// @Accept json
// @Produce json
// @Param body body dto.ClassDto true "ClassDto"
// @Success 201 {object} domain.Class
// @Failure 500 {object} error
// @Router /class [post]
func (h *ClassHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempClass dto.ClassDto
		if err := ctx.BindJSON(&tempClass); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		createdClass, err := h.service.Create(tempClass)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdClass)
	}
}

// class godoc
// @Summary Get all classes
// @Tags class
// @Produce json
// @Success 200 {array} domain.Class
// @Failure 500 {object} error
// @Router /class [get]
func (h *ClassHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		classList, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, classList)
	}
}

// class godoc
// @Summary Get class by id
// @Tags class
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Class
// @Failure 500 {object} error
// @Router /class/{id} [get]
func (h *ClassHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		class, err := h.service.GetById(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, class)
	}
}

// class godoc
// @Summary Update class
// @Tags class
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.ClassDto true "ClassDto"
// @Success 200 {object} domain.Class
// @Failure 500 {object} error
// @Router /class/{id} [put]
func (h *ClassHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var tempClass dto.ClassDto
		if err := ctx.BindJSON(&tempClass); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		updatedClass, err := h.service.Update(tempClass, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, updatedClass)
	}
}

// class godoc
// @Summary Delete class
// @Tags class
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /class/{id} [delete]
func (h *ClassHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		if err = h.service.Delete(intId); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, "Deleted Class with id: "+id)
	}
}
