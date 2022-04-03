package test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGroup(t *testing.T) {
	 app := fiber.New()

	 api := app.Group("/api")

	 v1 := api.Group("/v1")
	 v1.Get("/list", func(c *fiber.Ctx) error {
		 return c.SendString("List v1")
	 })
	 v1.Get("/user", func(c *fiber.Ctx) error {
		 return c.SendString("User v1")
	 })

	 v2 := api.Group("/v2")
	 v2.Get("/list", func(c *fiber.Ctx) error {
		 return c.SendString("List v2")
	 })
	 v2.Get("/user", func(c *fiber.Ctx) error {
		 return c.SendString("User v2")
	 })

	 app.Listen(":80")
}