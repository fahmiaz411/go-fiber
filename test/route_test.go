package test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRoute(t *testing.T) {
    app := fiber.New()

	app.Route("/test", func(router fiber.Router) {
		router.Get("/foo", func(c *fiber.Ctx) error {
			return c.SendString("foo")
		}).Name("foo")
		router.Get("/bar", func(c *fiber.Ctx) error {
			return c.SendString("bar")
		}).Name("bar")
	}, "test.")
}