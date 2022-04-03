package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName: "Fiber GO",
	})

	app2 := fiber.New()
	app2.Get("/status", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Mount("/v2", app2)

	Use(app)
	StaticFile(app)

	RouteGet(app)
	RoutePost(app)

	app.Server().MaxConnsPerIP = 1
	app.Listen("localhost:80")
	
	fmt.Println(app.HandlersCount())
}