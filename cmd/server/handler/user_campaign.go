package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/user_campaign"
)

type UserCampaignHandler struct {
	service user_campaign.UserCampaignService
}

func NewUserCampaignHandler(service *user_campaign.UserCampaignService) *UserCampaignHandler {
	return &UserCampaignHandler{service: *service}
}

// user_campaign godoc
// @Summary Create user_campaign
// @Tags user_campaign
// @Accept json
// @Produce json
// @Param body body dto.CreateUserCampaignDto true "CreateUserCampaignDto"
// @Success 201 {object} domain.UserCampaign
// @Failure 500 {object} error
// @Router /user_campaign [post]
func (h *UserCampaignHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempUserCampaign dto.CreateUserCampaignDto
		if err := ctx.BindJSON(&tempUserCampaign); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdUserCampaign, err := h.service.CreateUserCampaign(tempUserCampaign)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdUserCampaign)
	}
}

// user_campaign godoc
// @Summary Get all user_campaign
// @Tags user_campaign
// @Produce json
// @Success 200 {array} domain.UserCampaign
// @Failure 500 {object} error
// @Router /user_campaign [get]
func (h *UserCampaignHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userCampaignList, err := h.service.GetAllUserCampaigns()
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, userCampaignList)
	}
}

// user_campaign godoc
// @Summary Get user_campaign by id
// @Tags user_campaign
// @Produce json
// @Param id path int true "id"
// @Success 200 {array} domain.UserCampaign
// @Failure 500 {object} error
// @Router /user_campaign/{id} [get]
func (h *UserCampaignHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempUserCampaign, err := h.service.GetUserCampaignById(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, tempUserCampaign)
	}
}

// user_campaign godoc
// @Summary Get user_campaign by campaign id
// @Tags user_campaign
// @Produce json
// @Param id path int true "campaign_id"
// @Success 200 {array} domain.UserCampaign
// @Failure 500 {object} error
// @Router /user_campaign/campaign/{id} [get]
func (h *UserCampaignHandler) HandlerGetByCampaignId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempUserCampaign, err := h.service.GetUserCampaignByCampaignId(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, tempUserCampaign)
	}
}

// user_campaign godoc
// @Summary Get user_campaign by user id
// @Tags user_campaign
// @Produce json
// @Param id path int true "user_id"
// @Success 200 {array} domain.UserCampaign
// @Failure 500 {object} error
// @Router /user_campaign/user/{id} [get]
func (h *UserCampaignHandler) HandlerGetByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		tempUserCampaign, err := h.service.GetUserCampaignByUserId(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, tempUserCampaign)
	}
}

// user_campaign godoc
// @Summary Delete user_campaign
// @Tags user_campaign
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /user_campaign/{id} [delete]
func (h *UserCampaignHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		err = h.service.DeleteUserCampaign(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, "User Campaign deleted")
	}
}

func (h *UserCampaignHandler) HandlerAddFriendsToCampaign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var userIds []string
		if err := ctx.BindJSON(&userIds); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		err = h.service.AddFriendsToUserCampaign(userIds, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, "Friends added successfully")
	}
}

func (h *UserCampaignHandler) HandlerAddCharacterToCampaign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			log.Println(1, err.Error())
			ctx.JSON(400, err.Error())
			return
		}
		
		var body dto.AddCharacterToCampaignDto 
		if err := ctx.BindJSON(&body); err != nil {
			log.Println(2, err.Error())
			ctx.JSON(500, err.Error())
			return
		}
		
		
		err = h.service.AddCharacterToCampaign(body.CharacterId, body.CampaignId, cookie.Value)
		if err != nil {
			log.Println(3, err.Error())
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, "Character added successfully")
	}
}