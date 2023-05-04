package oops

import (
	"net/http"
	"strings"

	"github.com/go-errors/errors"
	fiber "github.com/gofiber/fiber/v2"
	arrayfuncs "github.com/izacgaldino23/array-funcs"
	"github.com/izacgaldino23/products-api/utils"
)

func Wrap(err error, msg string) error {
	return errors.WrapPrefix(err, msg, 0)
}

func Err(err error) error {
	return errors.Wrap(err, 1)
}

func HandleError(ctx *fiber.Ctx, err error) error {
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
			err = errors.Wrap(err, 0)
			stack := err.(*errors.Error).ErrorStack()

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
		ctx.Status(code).JSON(errorMessage)
		return err
	}

	// Return from handler
	return err
}
