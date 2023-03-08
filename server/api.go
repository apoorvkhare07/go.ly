package server

import "github.com/gofiber/fiber/v2"

func helloWorldApi(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
