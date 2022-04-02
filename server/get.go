package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RouteGet(r *fiber.App) {

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(strings.Repeat("H", 1))
	})

	r.Get("/:v", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("v"))
	})

	r.Get("/v/:n?", func(c *fiber.Ctx) error {
		n := c.Params("n")
		if n != "" {
			return c.SendString("Hello " + n)
		}

		return c.SendString("Where john ?")
	})

	r.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API Path: " + c.Params("*"))
	})


	r.Get("/file/*", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Custom not found")
	})
}