package main

import (
	"github.com/aditwibisana/golang-practice/auth-api/routes"

	"github.com/aditwibisana/golang-practice/auth-api/database"

	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(8000)
}
