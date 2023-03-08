package product

import (
	"net/http"

	goerrors "github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	app "github.com/izacgaldino23/products-api/application/product"
	"github.com/izacgaldino23/products-api/utils"
)

func AddProduct(c *fiber.Ctx) (err error) {
	var (
		product = &app.Product{}
	)

	err = c.BodyParser(product)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	id, err := app.AddProduct(product)
	if err != nil {
		return
	}

	return c.JSON(utils.T{
		"id": id,
	})
}

func ListProducts(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"status": http.StatusOK,
	})
}
