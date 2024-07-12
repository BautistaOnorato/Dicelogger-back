package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	skillxcharacterdata "github.com/proyecto-dnd/backend/internal/skillXCharacterData"
)

type SkillXCharacterHandler struct {
	service skillxcharacterdata.ServiceSkillXCharacter
}

func NewSkillXCharacterHandler(service *skillxcharacterdata.ServiceSkillXCharacter) *SkillXCharacterHandler {
	return &SkillXCharacterHandler{service: *service}
}

func (h *SkillXCharacterHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempSkillXCharacter domain.SkillXCharacterData
		if err := ctx.BindJSON(&tempSkillXCharacter); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdSkillXCharacter, err := h.service.Create(tempSkillXCharacter)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdSkillXCharacter)
	}
}

func (h *SkillXCharacterHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempSkillXCharacter domain.SkillXCharacterData
		if err := ctx.BindJSON(&tempSkillXCharacter); err != nil {
			ctx.JSON(500, err)
			return
		}
		err := h.service.Delete(tempSkillXCharacter)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, "Successfully deleted row with skill_id="+strconv.Itoa(int(tempSkillXCharacter.SkillID))+" and character_id="+strconv.Itoa(int(tempSkillXCharacter.CharacterID)))
	}
}

