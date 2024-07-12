package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	tradeevent "github.com/proyecto-dnd/backend/internal/tradeEvent"
)

type TradeEventHandler struct {
	service tradeevent.ServiceTradeEvent
}

func NewTradeEventHandler(service *tradeevent.ServiceTradeEvent) *TradeEventHandler {
	return &TradeEventHandler{service: *service}
}

func (h *TradeEventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempTradeEvent domain.TradeEvent
		if err := ctx.BindJSON(&tempTradeEvent); err != nil {
			ctx.JSON(400, err)
			return
		}
		createdTradeEvent, err := h.service.Create(tempTradeEvent)
		if err == tradeevent.ErrCannotBeNegative {
			ctx.JSON(400, err)
			return
		}
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdTradeEvent)
	}
}

func (h *TradeEventHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, err)
			return
		} 
		err = h.service.Delete(id)
		if err != nil {
			ctx.JSON(404, err)
			return	
		}
		ctx.JSON(204, nil)
	}
}

func (h *TradeEventHandler) HandlerGetBySessionId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		tradeEventList, err := h.service.GetBySessionId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tradeEventList)
	}
}

func (h *TradeEventHandler) HandlerGetBySender() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		tradeEventList, err := h.service.GetBySender(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tradeEventList)
	}
}

func (h *TradeEventHandler) HandlerGetByReceiver() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		tradeEventList, err := h.service.GetByReceiver(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tradeEventList)
	}
}

func (h *TradeEventHandler) HandlerGetBySenderOrReceiver() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		tradeEventList, err := h.service.GetBySenderOrReciever(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tradeEventList)
	}
}