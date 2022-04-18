package main

import (
	"C:/belajar_golang/src/auth-api/database"
	"C:/belajar_golang/src/auth-api/routes"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(8000)
}
