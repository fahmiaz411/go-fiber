package server

import (
	"github.com/gofiber/fiber/v2"
)

func StaticFile(r *fiber.App) {	
	r.Static("/file", "public")
}