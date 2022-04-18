package routes

import (
	"github.com/gofiber/fiber"
	"C:/belajar_golang/src/auth-api//controllers"
)

func setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
