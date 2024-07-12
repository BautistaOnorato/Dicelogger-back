package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/proficiency"
)

type ProficiencyHandler struct {
	service proficiency.ProficiencyService
}

func NewProficiencyHandler(service *proficiency.ProficiencyService) *ProficiencyHandler {
	return &ProficiencyHandler{service: *service}
}

// proficiency godoc
// @Summary Create proficiency
// @Tags proficiency
// @Accept json
// @Produce json
// @Param body body dto.ProficiencyDto true "ProficiencyDto"
// @Success 201 {object} domain.Proficiency
// @Failure 500 {object} error
// @Router /proficiency [post]
func (h *ProficiencyHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempProficiency dto.ProficiencyDto
		if err := ctx.BindJSON(&tempProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdProficiency, err := h.service.Create(tempProficiency)
		if err != nil {
			ctx.JSON(500, err)
		}
		ctx.JSON(201, createdProficiency)
	}
}

// proficiency godoc
// @Summary Get all proficiencies
// @Tags proficiency
// @Produce json
// @Success 200 {array} domain.Proficiency
// @Failure 500 {object} error
// @Router /proficiency [get]
func (h *ProficiencyHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proficiencyList, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, proficiencyList)
	}
}

// proficiency godoc
// @Summary Get proficiency by id
// @Tags proficiency
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Proficiency
// @Failure 500 {object} error
// @Router /proficiency/{id} [get]
func (h *ProficiencyHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		proficiency, err := h.service.GetById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, proficiency)
	}
}

func (h *ProficiencyHandler) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("characterId")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		proficiencyList, err := h.service.GetByCharacterDataId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, proficiencyList)
	}
}

// proficiency godoc
// @Summary Update proficiency
// @Tags proficiency
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.ProficiencyDto true "ProficiencyDto"
// @Success 200 {object} domain.Proficiency
// @Failure 500 {object} error
// @Router /proficiency/{id} [put]
func (h *ProficiencyHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		var tempProficiency dto.ProficiencyDto
		if err := ctx.BindJSON(&tempProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		updatedProficiency, err := h.service.Update(tempProficiency, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedProficiency)
	}
}

// proficiency godoc
// @Summary Delete proficiency
// @Tags proficiency
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
func (h *ProficiencyHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		if err = h.service.Delete(intId); err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Deleted Proficiency with id: "+id)
	}
}
