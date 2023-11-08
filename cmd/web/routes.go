package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"go-ws/internal/handler"
)

func routes() *fiber.App {
	engine := jet.New("./html", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// normal route
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
	app.Get("/", handler.Home)

	// ws
	app.Use(handler.Upgrade)
	app.Get("/ws", websocket.New(handler.ListenWs))

	return app
}
