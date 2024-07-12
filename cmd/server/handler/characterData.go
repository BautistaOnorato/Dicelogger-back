package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	characterdata "github.com/proyecto-dnd/backend/internal/characterData"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type CharacterHandler struct {
	service characterdata.ServiceCharacterData
}

func NewCharacterHandler(service *characterdata.ServiceCharacterData) *CharacterHandler {
	return &CharacterHandler{service: *service}
}

func (h *CharacterHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCharacterData domain.CharacterData
		if err := ctx.BindJSON(&tempCharacterData); err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		createdCharacterData, err := h.service.Create(tempCharacterData)
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(201, createdCharacterData)
	}
}

func (h *CharacterHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		err = h.service.Delete(id)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}

		ctx.JSON(204, nil)
	}
}

func (h *CharacterHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characters, err := h.service.GetAll()
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, characters)
	}
}

func (h *CharacterHandler) HandlerGetByCampaignIdAndUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		campaignid, errCampaign := strconv.Atoi(ctx.Query("campaignid"))
		userid := ctx.Query("userid")

		if errCampaign != nil && userid == "" {
			// temp error message
			ctx.JSON(400, errCampaign.Error())
			return
		}

		if errCampaign != nil {
			characters, err := h.service.GetByUserId(userid)
			if err != nil {
				ctx.JSON(500, err.Error())
				return
			}
			ctx.JSON(200, characters)
			return
		}

		if userid == "" {
			characters, err := h.service.GetByCampaignId(campaignid)
			if err != nil {
				ctx.JSON(500, err.Error())
				return
			}
			ctx.JSON(200, characters)
			return
		}

		characters, err := h.service.GetByUserIdAndCampaignId(userid, campaignid)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, characters)
	}
}

func (h *CharacterHandler) HandlerGetGenerics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characters, err := h.service.GetGenerics()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, characters)
	}
}

func (h *CharacterHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		characters, err := h.service.GetById(id)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, characters)
	}
}

func (h *CharacterHandler) HandlerGetByAttackEventId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventid, err := strconv.Atoi(ctx.Param("eventid"))
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		characters, err := h.service.GetByAttackEventId(eventid)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, characters)
	}
}

func (h *CharacterHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		var tempCharacterData domain.CharacterData
		if err := ctx.BindJSON(&tempCharacterData); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempCharacterData.Character_Id = id

		createdCharacterData, err := h.service.Update(tempCharacterData)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(201, createdCharacterData)
	}
}


func (h *CharacterHandler) HandlerGetByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		characters, err := h.service.GetByUser(cookie.Value)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, characters)
	}
}
