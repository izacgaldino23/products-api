package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/izacgaldino23/products-api/domain"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status": domain.GetTableName(&domain.PurchasedItem{}),
		})
	})

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal(err)
	}
}
