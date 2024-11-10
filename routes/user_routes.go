package routes

import (
	c "simpleapi/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", c.GetAllUsers)
	api.Get("/users/:id", c.GetUser)
	api.Post("/users", c.CreateUser)
	api.Put("/users/:id", c.UpdateUser)
	api.Delete("/users/:id", c.DeleteUser)
}
