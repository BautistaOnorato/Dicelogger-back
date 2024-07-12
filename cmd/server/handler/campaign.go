package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/campaign"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/user"
)

type CampaignHandler struct {
	service     campaign.CampaignService
	userService user.ServiceUsers
}

func NewCampaignHandler(service *campaign.CampaignService, userService *user.ServiceUsers) *CampaignHandler {
	return &CampaignHandler{service: *service, userService: *userService}
}

// campaign godoc
// @Summary Create campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Param body body dto.CreateCampaignDto true "CreateCampaignDto"
// @Success 201 {object} domain.Campaign
// @Failure 500 {object} error
// @Router /campaign [post]
func (h *CampaignHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		jwtClaimsInfo, err := h.userService.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		userId := jwtClaimsInfo.Id

		var tempCampaign dto.CreateCampaignDto
		if err := ctx.BindJSON(&tempCampaign); err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		createdCampaign, err := h.service.CreateCampaign(tempCampaign, userId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(201, createdCampaign)
	}
}

// campaign godoc
// @Summary Get all campaigns
// @Tags campaign
// @Produce json
// @Success 200 {array} dto.ResponseCampaignDto
// @Failure 500 {object} error
// @Router /campaign [get]
func (h *CampaignHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		campaignList, err := h.service.GetAllCampaigns()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, campaignList)
	}
}

// campaign godoc
// @Summary Get campaign by id
// @Tags campaign
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.ResponseCampaignDto
// @Failure 500 {object} error
// @Router /campaign/{id} [get]
func (h *CampaignHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		tempCampaign, err := h.service.GetCampaignByID(intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempCampaign)
	}
}

// campaign godoc
// @Summary Get campaigns by user id
// @Tags campaign
// @Produce json
// @Param id path int true "user_id"
// @Success 200 {array} dto.ResponseCampaignDto
// @Failure 500 {object} error
// @Router /campaign/user/{id} [get]
func (h *CampaignHandler) HandlerGetByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		tempCampaign, err := h.service.GetCampaignsByUserId(cookie.Value)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, tempCampaign)
	}
}

// campaign godoc
// @Summary Update campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.CreateCampaignDto true "CreateCampaignDto"
// @Success 200 {object} dto.ResponseCampaignDto
// @Failure 500 {object} error
// @Router /campaign/{id} [put]
func (h *CampaignHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		var tempCampaign dto.CreateCampaignDto
		if err := ctx.BindJSON(&tempCampaign); err != nil {

			ctx.JSON(500, err.Error())
			return
		}

		updatedCampaign, err := h.service.UpdateCampaign(tempCampaign, intId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, updatedCampaign)
	}
}

// campaign godoc
// @Summary Delete campaign
// @Tags campaign
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /campaign/{id} [delete]
func (h *CampaignHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		serviceErr := h.service.DeleteCampaign(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr.Error())
			return
		}

		ctx.JSON(200, "Deleted Campaign with id "+id)
	}
}

