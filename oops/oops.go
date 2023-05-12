package oops

import (
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/go-errors/errors"
	fiber "github.com/gofiber/fiber/v2"
	arrayfuncs "github.com/izacgaldino23/array-funcs"
	"github.com/izacgaldino23/products-api/utils"
)

var IsOnPanic bool

func Wrap(err error, msg string) error {
	return errors.WrapPrefix(err, msg, 0)
}

func Err(err error) error {
	return errors.Wrap(err, 1)
}

func HandleError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if IsOnPanic {
		IsOnPanic = false
		return err
	}

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
			err = errors.Wrap(err, 1)
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
		_ = ctx.Status(code).JSON(errorMessage)
		return err
	}

	// Return from handler
	return err
}

func HandleErrorRecovery(ctx *fiber.Ctx, err interface{}) {
	IsOnPanic = true
	var msg string

	switch e := err.(type) {
	case string:
		msg = e
	default:
		msg = err.(error).Error()
	}

	formatStack(msg)
}

func formatStack(err string) {
	var (
		stackColor        = arrayfuncs.AnyToArrayKind(strings.Split(string(debug.Stack()), "\n"))
		started, finished bool
	)

	stackColor.ForEach(func(line string, i int, a *[]string) {
		if i == 0 {
			println(utils.Colorize(utils.Purple, "IS ON PANIC: "+err))
			started = true
		}

		if started && !finished {
			if strings.Contains(line, "panic.go") {
				finished = true
			}

			return
		}

		if !strings.Contains(line, "	") {
			sentences := strings.Split(line, "(0x")

			sentences[0] = utils.Colorize(utils.Red, sentences[0])
			print(sentences[0], " ")

			if len(sentences) > 1 {
				sentences[1] = utils.Colorize(utils.Gray, "(0x"+sentences[1])
				println(sentences[1])
			} else {
				println()
			}

			// line = strings.Join(sentences, " ")
		} else if v := strings.TrimSpace(line); v != "" {
			println(utils.Colorize(utils.Green, "   ->"), v)
		}
	})
}
