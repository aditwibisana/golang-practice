package controllers

import "github.com/gofiber/fiber"

func Hello(c *fiber.Ctx) {
	c.Send("Hello Njink")
}
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(503).Send(err)
		return err
	}
	return c.JSON(data)
}
