package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	UserRoutes(app)
	ProductRoutes(app)
}