package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Use(r *fiber.App) {
	r.Use("/api/*", func(c *fiber.Ctx) error {
		fmt.Println("api middleware")
		return c.Next()
	})
}