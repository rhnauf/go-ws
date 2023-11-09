package handler

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-ws/internal/entity"
	"go-ws/internal/service"
	"log"
	"time"
)

type WsHandler struct {
	messageService service.MessageService
}

func NewWsHandler(messageService service.MessageService) *WsHandler {
	return &WsHandler{
		messageService: messageService,
	}
}

func (h *WsHandler) Home(ctx *fiber.Ctx) error {
	return ctx.Render("home", nil)
}

func (h *WsHandler) Upgrade(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {
		ctx.Locals("allowed", true)
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

var wsChan = make(chan WsRequest)

type WebSocketConnection struct {
	*websocket.Conn
}

type WsRequest struct {
	Action    string              `json:"action"`
	Username  string              `json:"username"`
	Recipient string              `json:"recipient"`
	Message   string              `json:"message"`
	Conn      WebSocketConnection `json:"-"`
}

func (h *WsHandler) ListenWs(ctx *websocket.Conn) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsRequest
	for {
		if err := ctx.Conn.ReadJSON(&payload); err != nil {
			log.Println(err)
		} else {
			payload.Conn = WebSocketConnection{ctx}
			log.Printf("request payload: %+v\n", payload)
			wsChan <- payload
		}
	}
}

type User struct {
	id       string
	username string
}

var clients = make(map[WebSocketConnection]User)

type WsResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

func (h *WsHandler) Hub() {
	var response WsResponse

	for {
		event := <-wsChan
		sender := event.Username
		recipient := event.Recipient
		switch event.Action {
		case "connect":
			clients[event.Conn] = User{
				id:       uuid.New().String(),
				username: event.Username,
			}
			response.Action = "list_users"
			findAndBroadcast(response, sender, recipient)
		case "left":
			response.Action = "list_users"
			delete(clients, event.Conn)
			findAndBroadcast(response, sender, recipient)
		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("%v &emsp; <strong>%s</strong>: %s", time.Now().Local().Format("2006-01-02 15:04:05"), event.Username, event.Message)

			// insert message to db
			m := &entity.Message{
				Id:        uuid.New().String(),
				Username:  event.Username,
				Recipient: event.Recipient,
				Message:   event.Message,
			}
			err := h.messageService.CreateMessage(m)
			if err != nil {
				log.Println(err)
				continue
			}

			log.Println("entry created successfully =>", m)

			findAndBroadcast(response, sender, recipient)
		}
	}
}

func findAndBroadcast(response WsResponse, sender, recipient string) {
	for conn := range clients {
		if clients[conn].username == sender || clients[conn].username == recipient {
			err := conn.Conn.WriteJSON(response)
			if err != nil {
				log.Println("error occurred broadcasting message")
				_ = conn.Conn.Close()
				delete(clients, conn)
			}
		}
	}
}
