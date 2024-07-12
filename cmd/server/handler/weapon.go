package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/weapon"
)

type WeaponHandler struct {
	service weapon.ServiceWeapon
}

func NewWeaponHandler(service *weapon.ServiceWeapon) *WeaponHandler{
	return &WeaponHandler{service: *service}
}

// weapon godoc
// @Summary Create weapon
// @Tags weapon
// @Accept json
// @Produce json
// @Param body body domain.Weapon true "Weapon"
// @Success 201 {object} domain.Weapon
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon [post]
func (h *WeaponHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeapon domain.Weapon
		if err := ctx.BindJSON(&tempWeapon); err != nil {
			ctx.AbortWithError(400, err)
            return
		}

		createdWeapon, err := h.service.Create(tempWeapon)
		if err!= nil {
            ctx.AbortWithError(500, err)
            return
        }
		ctx.JSON(201, createdWeapon)
	}
}

// weapon godoc
// @Summary Delete weapon
// @Tags weapon
// @Produce json
// @Param id path int true "id"
// @Success 204 {object} nil
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /weapon/{id} [delete]
func (h *WeaponHandler) HandlerDelete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		err = h.service.Delete(id)
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(204, nil)
	}
}

// weapon godoc
// @Summary Get all weapons
// @Tags weapon
// @Produce json
// @Success 200 {array} domain.Weapon
// @Failure 500 {object} error
// @Router /weapon [get]
func (h *WeaponHandler) HandlerGetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		weapons, err := h.service.GetAll()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapons)
	}
}

// weapon godoc
// @Summary Get weapons by campaign id
// @Tags weapon
// @Produce json
// @Param id path int true "campaign_id"
// @Success 200 {array} domain.Weapon
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /weapon/campaign/{id} [get]
func (h *WeaponHandler) HandlerGetByCampaignId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}
		weapons, err := h.  service.GetByCampaignId(id)
		if err == weapon.ErrNotFound {
			ctx.AbortWithError(404, err)
            return
		}
		if err != nil{
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapons)
	}
}

// weapon godoc
// @Summary Get weapon by id
// @Tags weapon
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Weapon
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon/{id} [get]
func (h *WeaponHandler) HandlerGetById() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}
		weapon, err := h.  service.GetById(id)
		if err != nil{
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapon)
	}
}

// weapon godoc
// @Summary Get all generic weapons
// @Tags weapon
// @Produce json
// @Success 200 {array} domain.Weapon
// @Failure 500 {object} error
// @Router /weapon/generic [get]
func (h *WeaponHandler) HandlerGetAllGeneric() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		weapons, err := h.service.GetAllGeneric()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapons)
	}
}

// weapon godoc
// @Summary Update weapon
// @Tags weapon
// @Accept json
// @Produce json
// @Param body body domain.Weapon true "Weapon"
// @Param id path int true "id"
// @Success 200 {object} domain.Weapon
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon/{id} [put]
func (h *WeaponHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeapon domain.Weapon
		if err := ctx.BindJSON(&tempWeapon); err != nil {
			ctx.AbortWithError(400, err)
            return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}

		tempWeapon.Weapon_Id = id

		updatedWeapon, err := h.service.Update(tempWeapon)
		if err!= nil {
            ctx.AbortWithError(500, err)
            return
        }
		ctx.JSON(200, updatedWeapon)
	}
}