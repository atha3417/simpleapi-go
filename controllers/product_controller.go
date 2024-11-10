package controllers

import (
	repo "simpleapi/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	products, err := repo.GetAllProducts();
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get users"})
	}
	return c.JSON(products)
}