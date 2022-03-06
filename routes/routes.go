package routes

import (
	"go-lang-react/controllers"

	"github.com/gofiber/fiber/v2"
)

func StartApplication(app *fiber.App) {
	app.Post("/register", controllers.Register)

}