package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
)

type WeaponXCharacterDataHandler struct {
	service weaponxcharacterdata.ServiceWeaponXCharacterData
}

func NewWeaponXCharacterDataHandler(service *weaponxcharacterdata.ServiceWeaponXCharacterData) *WeaponXCharacterDataHandler {
	return &WeaponXCharacterDataHandler{service: *service}
}

// weaponXCharacterData godoc
// @Summary Create weaponXCharacterData
// @Tags weaponXCharacterData
// @Accept json
// @Produce json
// @Param body body domain.WeaponXCharacterData true "WeaponXCharacterData"
// @Success 201 {object} domain.WeaponXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon_character [post]
func (h *WeaponXCharacterDataHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeaponXCharacterData domain.WeaponXCharacterData
		if err := ctx.BindJSON(&tempWeaponXCharacterData); err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		createdWeaponXCharacterData, err := h.service.Create(tempWeaponXCharacterData)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(201, createdWeaponXCharacterData)
	}
}

// weaponXCharacterData godoc
// @Summary Delete weaponXCharacterData
// @Tags weaponXCharacterData
// @Produce json
// @Param id path int true "id"
// @Success 204 {object} nil
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /weapon_character/{id} [delete]
func (h *WeaponXCharacterDataHandler) HandlerDelete() gin.HandlerFunc {
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

// weaponXCharacterData godoc
// @Summary Delete weaponXCharacterData by character id
// @Tags weaponXCharacterData
// @Produce json
// @Param id path int true "character_id"
// @Success 204 {object} nil
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /weapon_character/character/{id} [delete]
func (h *WeaponXCharacterDataHandler) HandlerDeleteByCharacterDataId() gin.HandlerFunc {
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

// weaponXCharacterData godoc
// @Summary Get all weaponXCharacterData
// @Tags weaponXCharacterData
// @Produce json
// @Success 200 {array} domain.WeaponXCharacterData
// @Failure 500 {object} error
// @Router /weapon_character [get]
func (h *WeaponXCharacterDataHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		weaponXCharacterDataList, err := h.service.GetAll()
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		ctx.JSON(200, weaponXCharacterDataList)
	}
}

// weaponXCharacterData godoc
// @Summary Get weaponXCharacterData by character id
// @Tags weaponXCharacterData
// @Produce json
// @Param id path int true "character_id"
// @Success 200 {array} domain.WeaponXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon_character/character/{id} [get]
func (h *WeaponXCharacterDataHandler) HandlerGetByCharacterDataId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		weaponXCharacterDataList, err := h.service.GetByCharacterDataId(id)
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(200, weaponXCharacterDataList)
	}
}

// weaponXCharacterData godoc
// @Summary Get weaponXCharacterData by id
// @Tags weaponXCharacterData
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.WeaponXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon_character/{id} [get]
func (h *WeaponXCharacterDataHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		weaponXCharacterData, err := h.service.GetById(id)
		if err != nil {
			// fmt.Println(err)
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(200, weaponXCharacterData)
	}
}

// weaponXCharacterData godoc
// @Summary Update weaponXCharacterData
// @Tags weaponXCharacterData
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body domain.WeaponXCharacterData true "WeaponXCharacterData"
// @Success 200 {object} domain.WeaponXCharacterData
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /weapon_character/{id} [put]
func (h *WeaponXCharacterDataHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeaponXCharacterData domain.WeaponXCharacterData
		if err := ctx.BindJSON(&tempWeaponXCharacterData); err != nil {
			ctx.AbortWithError(400, err)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}

		tempWeaponXCharacterData.Character_Weapon_Id = id

		updatedWeaponXCharacterData, err := h.service.Update(tempWeaponXCharacterData)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(200, updatedWeaponXCharacterData)
	}
}
