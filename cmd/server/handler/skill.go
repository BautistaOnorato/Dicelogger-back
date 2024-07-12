package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/skill"
)

type SkillHandler struct {
	service skill.ServiceSkill
}

func NewSkillHandler(service *skill.ServiceSkill) *SkillHandler {
	return &SkillHandler{service: *service}
}

func (h *SkillHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempSkill domain.Skill
		if err := ctx.BindJSON(&tempSkill); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdSkill, err := h.service.Create(tempSkill)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdSkill)
	}
}

func (h *SkillHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		err = h.service.Delete(intId)
		if err != nil {
			ctx.JSON(404, err)
			return
		}
		ctx.JSON(200, "Deleted skill with id: "+id)
	}
}

func (h *SkillHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skills, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, skills)
	}
}
func (h *SkillHandler) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		skills, err := h.service.GetByCharacterId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, skills)
	}
}

func (h *SkillHandler) HandlerGetByClassId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		skills, err := h.service.GetByClassId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, skills)
	}
}

func (h *SkillHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		skill, err := h.service.GetById(intId)
		if err != nil {
			ctx.JSON(404, err)
			return
		}
		ctx.JSON(200, skill)
	}
}

func (h *SkillHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		var tempSkill domain.Skill
		if err := ctx.BindJSON(&tempSkill); err != nil {
			ctx.JSON(500, err)
			return
		}
		tempSkill.SkillId = id
		updatedSkill, err := h.service.Update(tempSkill)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedSkill)
	}
}
