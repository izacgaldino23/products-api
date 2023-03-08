package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	goerrors "github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2"
	arrayfuncs "github.com/izacgaldino23/array-funcs"
	"github.com/izacgaldino23/products-api/config"
	"github.com/izacgaldino23/products-api/handler/product"
	"github.com/izacgaldino23/products-api/utils"
)

func init() {
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	err = config.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: handleError,
	})

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

func handleError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if err != nil {
		errorMessage := utils.T{
			"error": err.Error(),
		}

		if code != http.StatusUnprocessableEntity {
			stack := err.(*goerrors.Error).ErrorStack()

			stackColor := arrayfuncs.AnyToArrayKind(strings.Split(stack, "\n"))

			stackColor.ForEach(func(line string, i int, a *[]string) {
				if strings.Contains(line, " (") {
					sentences := strings.Split(line, " ")

					sentences[0] = utils.Colorize(utils.Red, sentences[0])
					sentences[1] = utils.Colorize(utils.Gray, sentences[1])

					print(sentences[0], " ")
					println(sentences[1])

					// line = strings.Join(sentences, " ")
				} else if v := strings.TrimSpace(line); v != "" {
					println(utils.Colorize(utils.Green, "   ->"), v)
				}
			})

			stack = strings.ReplaceAll(stack, "\t", "    ")
			stack = strings.ReplaceAll(stack, "\n", "\n    ")

			errorMessage["stack"] = strings.Split(stack, "\n")
		}

		// In case the SendFile fails
		return ctx.Status(code).JSON(errorMessage)
	}

	// Return from handler
	return nil
}
