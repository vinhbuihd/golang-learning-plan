package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/day14-shopi-api/handlers"
	"github.com/yourname/day14-shopi-api/models"
)

var products = []models.Product{
	{ID: 1, Name: "Product 1", Price: 10.0},
	{ID: 2, Name: "Product 2", Price: 20.0},
	{ID: 3, Name: "Product 3", Price: 30.0},
}

func main() {
	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Health check passed")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Shop API")
	})

	app.Get("/products", handlers.GetProducts)
	app.Get("/products/:id", handlers.GetProduct)
	app.Post("/products", handlers.CreateProduct)
	app.Put("/products/:id", handlers.UpdateProduct)
	app.Delete("/products/:id", handlers.DeleteProduct)

	app.Listen(":3000")

}
