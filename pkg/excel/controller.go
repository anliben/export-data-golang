package excel

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	routes := app.Group("/excel")
	routes.Post("/", Export)
}
