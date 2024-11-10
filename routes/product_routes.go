package routes

import (
	c "simpleapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", c.GetAllProducts)
}