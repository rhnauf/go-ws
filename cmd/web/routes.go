package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"github.com/jmoiron/sqlx"
	"go-ws/internal/handler"
	"go-ws/internal/repository"
	"go-ws/internal/service"
	"log"
)

func routes(dbConn *sqlx.DB) *fiber.App {
	var (
		engine = jet.New("./html", ".jet")
		app    = fiber.New(fiber.Config{
			Views: engine,
		})
		messageRepository = repository.NewMessageRepository(dbConn)
		messageService    = service.NewMessageService(messageRepository)
		messageHandler    = handler.NewMessageHandler(messageService)
		wsHandler         = handler.NewWsHandler(messageService)
	)

	// normal route
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/message/list", messageHandler.GetListMessage)
	app.Get("/", wsHandler.Home)

	// ws
	app.Use(wsHandler.Upgrade)

	// initialize chat hub
	log.Println("starting chat hub")
	go wsHandler.Hub()

	app.Get("/ws", websocket.New(wsHandler.ListenWs))

	return app
}
