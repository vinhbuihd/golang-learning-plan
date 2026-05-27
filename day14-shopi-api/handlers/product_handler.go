package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yourname/day14-shopi-api/models"
)

var products = []models.Product{
	{ID: 1, Name: "Keyboard", Price: 120},
	{ID: 2, Name: "Mouse", Price: 50},
	{ID: 3, Name: "Monitor", Price: 300},
}

func GetProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	for _, product := range products {
		if fmt.Sprint(product.ID) == id {
			return c.JSON(product)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Product not found",
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	products = append(products, product)

	return c.Status(201).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedProduct models.Product

	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	for index, product := range products {
		if fmt.Sprint(product.ID) == id {
			products[index].Name = updatedProduct.Name
			products[index].Price = updatedProduct.Price

			return c.JSON(products[index])
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Product not found",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	for index, product := range products {
		if fmt.Sprint(product.ID) == id {
			products = append(products[:index], products[index+1:]...)

			return c.JSON(fiber.Map{
				"message": "Product deleted",
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Product not found",
	})
}
