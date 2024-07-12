package ws

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/attackEvent"
	"github.com/proyecto-dnd/backend/internal/dice_event"
	tradeevent "github.com/proyecto-dnd/backend/internal/tradeEvent"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
	tradeEventService tradeevent.ServiceTradeEvent
	attackEventService attackEvent.AttackEventService
	diceEventService dice_event.DiceEventService
}

func NewHub(tradeEventService tradeevent.ServiceTradeEvent, attackEventService attackEvent.AttackEventService, diceEventService dice_event.DiceEventService) *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		tradeEventService: tradeEventService,
		attackEventService: attackEventService,
		diceEventService: diceEventService,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				fmt.Println("Unregistered client: ", client)
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if message.SessionID == client.sessionId {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) ServeWs(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	
	sessionId, err := strconv.Atoi(ctx.Param("session_id"))
	if err != nil {
		log.Println(err)
		return
	}
	
	client := NewClient(h, conn, sessionId)
	h.register <- client

	log.Println("Nuevo cliente conectado: ", client.sessionId)

	go client.readPump()
	go client.writePump()
}