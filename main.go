package main

import (
	// "./database"
	"github.com/gofiber/fiber/v2"
	"github.com/winsongr/go-lang-react/routes"
)

func main() {
	// database.Connect()
	app := fiber.New()
	routes.StartApplication(app)
	app.Listen(":8000")
}
