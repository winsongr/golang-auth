package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/winsongr/go-lang-react/controllers"
)

func StartApplication(app *fiber.App) {
	app.Get("/", controllers.Home)

}
