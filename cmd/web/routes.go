package main

import "github.com/gofiber/fiber/v2"

func routes() *fiber.App {
	mux := fiber.New()

	mux.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	return mux
}
