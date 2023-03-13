package product

import (
	goerrors "github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	app "github.com/izacgaldino23/products-api/application/product"
	"github.com/izacgaldino23/products-api/utils"
)

func AddProduct(c *fiber.Ctx) (err error) {
	var (
		product = &app.Product{}
	)

	err = utils.Validate(product, c)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	id, err := app.AddProduct(product)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	return c.JSON(utils.T{
		"id": id,
	})
}

func ListProducts(c *fiber.Ctx) error {
	params := utils.ParseParams(c)

	products, err := app.ListProducts(&params)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	return c.JSON(products)
}
