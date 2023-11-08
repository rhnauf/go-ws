package handler

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"sort"
	"time"
)

func Home(ctx *fiber.Ctx) error {
	return ctx.Render("home", nil)
}

func Upgrade(ctx *fiber.Ctx) error {
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
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

func ListenWs(ctx *websocket.Conn) {
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
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

func Hub() {
	var response WsResponse

	for {
		event := <-wsChan
		switch event.Action {
		case "connect":
			clients[event.Conn] = User{
				id:       uuid.New().String(),
				username: event.Username,
			}
			response.Action = "list_users"
			response.ConnectedUsers = getUserList()
			broadcastAll(response)
		case "left":
			response.Action = "list_users"
			delete(clients, event.Conn)
			response.ConnectedUsers = getUserList()
			broadcastAll(response)
		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("%v &emsp; <strong>%s</strong>: %s", time.Now().Local().Format("2006-01-02 15:04:05"), event.Username, event.Message)
			broadcastAll(response)
		}
	}
}

func getUserList() []string {
	var userList []string
	for _, conn := range clients {
		if conn.username != "" {
			userList = append(userList, conn.username)
		}
	}
	sort.Strings(userList)
	return userList
}

func broadcastAll(response WsResponse) {
	for conn := range clients {
		err := conn.Conn.WriteJSON(response)
		if err != nil {
			log.Println("error occurred broadcasting message")
			_ = conn.Conn.Close()
			delete(clients, conn)
		}
	}
}
