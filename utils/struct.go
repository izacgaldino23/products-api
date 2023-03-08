package utils

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	errTagIsEmpty = errors.New("The field doesn't have Tag")
)

type T map[string]interface{}

func ParseStructToMap(m interface{}) (generatedMap map[string]interface{}, err error) {
	var (
		value         = reflect.ValueOf(m)
		typeOf        = reflect.TypeOf(m)
		tagName       = "sql"
		tagIgnoreName = "ignoreInsertUpdate"
	)

	generatedMap = make(map[string]interface{})

	for i := 0; i < typeOf.Elem().NumField(); i++ {
		fieldName := typeOf.Elem().Field(i).Tag.Get(tagName)
		ignore := typeOf.Elem().Field(i).Tag.Get(tagIgnoreName)

		if fieldName == "" {
			err = errTagIsEmpty
			return
		} else if fieldName == "-" || ignore != "" {
			continue
		}

		generatedMap[fieldName] = value.Elem().Field(i).Interface()
	}

	return
}

func Validate[T any](s *T, c *fiber.Ctx) (err error) {
	_ = c.BodyParser(s)

	err = validator.New().Struct(s)
	if err != nil {
		msg := ""
		for _, e := range err.(validator.ValidationErrors) {
			// el.Field = err.Field()
			// el.Tag = err.Tag()
			// el.Value = err.Param()
			msg += fmt.Sprintf("\nField %v is %v", e.Field(), e.Tag())
			// fiber.err
		}

		err = fiber.NewError(http.StatusUnprocessableEntity, msg)

		return
	}

	return
}
