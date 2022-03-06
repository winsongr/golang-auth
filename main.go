package main

import (
	// "./database"
	"go-lang-react/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// database.Connect()
	app := fiber.New()
	routes.StartApplication(app)
	app.Listen(":8000")
}
