package routes

import (
	"github.com/aditwibisana/golang-practice/auth-api/controllers"

	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Get("/hallo", controllers.Hello)
	app.Get("/register", func(c *fiber.Ctx) {
		var data map[string]string
		if err := c.BodyParser(&data); err != nil {
			c.Status(503).Send(err)
			return
		}
		c.JSON(data)
	})
}
