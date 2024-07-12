package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/proyecto-dnd/backend/internal/domain"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 10 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 2048
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub       *Hub
	conn      *websocket.Conn
	send      chan *Message
	sessionId int
}

type Message struct {
	Content   EventData `json:"content"`
	Sent      time.Time `json:"sent"`
	SessionID int       `json:"session_id"`
}

func NewClient(hub *Hub, conn *websocket.Conn, sessionId int) *Client {
	return &Client{
		hub:       hub,
		conn:      conn,
		send:      make(chan *Message),
		sessionId: sessionId,
	}
}

// func (c *Client) pongHandler(pongMsg string) error {
// 	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
// }

type EventData struct {
	Type      string          `json:"type"`
	EventData json.RawMessage `json:"eventData"`
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	// if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// c.conn.SetPongHandler(c.pongHandler)

	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var event EventData

		err = json.Unmarshal(message, &event)
		if err != nil || event.Type == "" || event.EventData == nil {
			log.Printf("error: %v", err)
			continue
		}

		switch event.Type {
		case "trade":
			
		case "attack":
			var attackEventDto domain.AttackEvent
			err = json.Unmarshal(event.EventData, &attackEventDto)
			if err != nil {
				log.Printf("error: %v", err)
				continue
			} else {
				attackEvent, err := c.hub.attackEventService.CreateEvent(attackEventDto)
				if err != nil {
					log.Printf("error: %v", err)
					continue
				}
				event.EventData, _ = json.Marshal(attackEvent)
			}
		case "dice":
			var diceEvent domain.DiceEvent
			err = json.Unmarshal(event.EventData, &diceEvent)
			if err != nil {
				log.Printf("error: %v", err)
				continue
			} else {
				diceEvent, err = c.hub.diceEventService.Create(diceEvent)
				if err != nil {
					log.Printf("error: %v", err)
					continue
				}
				event.EventData, _ = json.Marshal(diceEvent)
			}
		}

		msg := &Message{
			Content:   event,
			Sent:      time.Now(),
			SessionID: c.sessionId,
		}

		log.Println(msg)
		log.Println(event)

		c.hub.broadcast <- msg
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))

			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				return
			}

			w.Write([]byte(jsonMessage))

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
