package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/izacgaldino23/products-api/config"
	"github.com/izacgaldino23/products-api/config/database"
	"github.com/izacgaldino23/products-api/handler/product"
	"github.com/izacgaldino23/products-api/oops"
)

func init() {
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	err = database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: oops.HandleError,
	})

	app.Use(recover.New(recover.Config{
		StackTraceHandler: oops.HandleErrorRecovery,
		EnableStackTrace:  true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"status": 200,
		})
	})

	handle(app)

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal(err)
	}
}

func handle(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	product.Router(v1.Group("/product"))
}
