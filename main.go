package main

import (
		"go-lang-react/routes"
		"go-lang-react/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.StartApplication(app)
	app.Listen(":8000")
}
