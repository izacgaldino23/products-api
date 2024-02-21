package product

import (
	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	r.Post("/", AddProduct)
	r.Get("/", ListProducts)
	r.Get("/:product_id", GetProduct)
	r.Put("/:product_id", UpdateProduct)
	r.Delete("/:product_id", DeleteProduct)
}
