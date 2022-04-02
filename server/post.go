package server

import "github.com/gofiber/fiber/v2"

func RoutePost(r *fiber.App){

	r.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("Im post")
	})
}