package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/item"
)

type ItemHandler struct {
	service item.ServiceItem
}

func NewItemHandler(service *item.ServiceItem) *ItemHandler {
    return &ItemHandler{service: *service}
}

// item godoc
// @Summary Create item
// @Tags item
// @Accept json
// @Produce json
// @Param body body domain.Item true "Item"
// @Success 201 {object} domain.Item
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /item [post]
func (h *ItemHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempItem domain.Item
        if err := ctx.BindJSON(&tempItem); err!= nil {
            ctx.JSON(400, err)
            return
        }
		createdItem, err := h.service.Create(tempItem)
		if err!= nil {
            ctx.JSON(500, err)
            return
        }
		ctx.JSON(201, createdItem)
    }
}

// item godoc
// @Summary Delete item
// @Tags item
// @Produce json
// @Param id path int true "id"
// @Success 204 {object} nil
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Router /item/{id} [delete]
func (h *ItemHandler) HandlerDelete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
        id, err := strconv.Atoi(ctx.Param("id"))
        if err!= nil {
            ctx.JSON(400, err)
            return
        }
        err = h.service.Delete(id)
        if err!= nil {
            ctx.JSON(404, err)
            return
        }
        ctx.JSON(204, nil)
    }
}

// item godoc
// @Summary Get all items
// @Tags item
// @Produce json
// @Success 200 {array} domain.Item
// @Failure 500 {object} error
// @Router /item [get]
func (h *ItemHandler) HandlerGetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		items, err := h.service.GetAll()
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        ctx.JSON(200, items)
	}
}

// item godoc
// @Summary Get items by campaign id
// @Tags item
// @Produce json
// @Param id path int true "campaign_id"
// @Success 200 {array} domain.Item
// @Failure 500 {object} error
// @Router /item/campaign/{id} [get]
func (h *ItemHandler) HandlerGetByCampaignId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
        id := ctx.Param("id")

        intId, err := strconv.Atoi(id)
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        items, err := h.service.GetByCampaignId(intId)
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        ctx.JSON(200, items)
    }
}

// item godoc
// @Summary Get item by id
// @Tags item
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Item
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /item/{id} [get]
func (h *ItemHandler) HandlerGetById() gin.HandlerFunc{
	return func(ctx *gin.Context) {
        id := ctx.Param("id")

        intId, err := strconv.Atoi(id)
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        items, err := h.service.GetById(intId)
        if err!= nil {
            ctx.JSON(404, err)
            return
        }
        ctx.JSON(200, items)
    }
}

// item godoc
// @Summary Get all generic items
// @Tags item
// @Produce json
// @Success 200 {array} domain.Item
// @Failure 500 {object} error
// @Router /item/generic [get]
func (h *ItemHandler) HandlerGetAllGeneric() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		items, err := h.service.GetAllGeneric()
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        ctx.JSON(200, items)
	}
}

// item godoc
// @Summary Update item
// @Tags item
// @Accept json
// @Produce json
// @Param body body domain.Item true "Item"
// @Param id path int true "id"
// @Success 200 {object} domain.Item
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /item/{id} [put]
func (h *ItemHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempItem domain.Item
        if err := ctx.BindJSON(&tempItem); err!= nil {
            ctx.JSON(400, err)
            return
        }

        id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}

        tempItem.Item_Id = id

		updatedItem, err := h.service.Update(tempItem)
		if err!= nil {
            ctx.JSON(500, err)
            return
        }
		ctx.JSON(200, updatedItem)
    }
}