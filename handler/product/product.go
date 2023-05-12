package product

import (
	"strconv"

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

	return c.Status(201).
		JSON(utils.T{
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

func UpdateProduct(c *fiber.Ctx) (err error) {
	var (
		product = &app.Product{}
		id      int64
	)

	err = utils.Validate(product, c)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	id, err = strconv.ParseInt(c.Params("product_id"), 10, 64)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	err = app.UpdateProduct(id, product)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	return c.SendStatus(204)
}

func DeleteProduct(c *fiber.Ctx) (err error) {
	var (
		id int64
	)

	id, err = strconv.ParseInt(c.Params("product_id"), 10, 64)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	err = app.DeleteProduct(id)
	if err != nil {
		return goerrors.Wrap(err, 0)
	}

	return c.SendStatus(204)
}
