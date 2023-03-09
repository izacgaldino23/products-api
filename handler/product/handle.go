package product

import (
	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	r.Post("/", AddProduct)
	r.Get("/", ListProducts)
	// r.Post("/", UpdateProduct)
}
