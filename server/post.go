package server

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
	User string `json:"user"`
}

func RoutePost(r *fiber.App){

	r.Post("/api/register", func(c *fiber.Ctx) error {

		body := new(Body)
		c.BodyParser(body)

		c.Set("Content-type", "application/json")
		res, _ := json.Marshal(map[string]interface{}{
			"message": "Hello " + body.User,
		})
		return c.Send(res)
	})
}