package controllers

import "github.com/gofiber/fiber"

func hello(c *fiber.Ctx) error {
	c.SendString("Hello, World!")
}
