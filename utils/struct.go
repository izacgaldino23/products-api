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
		} else if fieldName == "-" || ignore != "" || value.Elem().Field(i).IsNil() {
			continue
		}

		generatedMap[fieldName] = value.Elem().Field(i).Interface()
	}

	return
}

func GetSqlColumnList(m interface{}) (columns []string, err error) {
	var (
		typeOf  = reflect.TypeOf(m)
		tagName = "sql"
	)

	columns = make([]string, 0)

	for i := 0; i < typeOf.Elem().NumField(); i++ {
		fieldName := typeOf.Elem().Field(i).Tag.Get(tagName)

		if fieldName == "" {
			err = errTagIsEmpty
			return
		} else if fieldName == "-" {
			continue
		}

		columns = append(columns, fieldName)
	}

	return
}

func Validate[T any](s *T, c *fiber.Ctx) (err error) {
	_ = c.BodyParser(s)

	err = validator.New().Struct(s)
	if err != nil {
		msg := ""
		for _, e := range err.(validator.ValidationErrors) {

			msg += fmt.Sprintf("\nField %v is %v", e.Field(), e.Tag())

		}

		err = fiber.NewError(http.StatusUnprocessableEntity, msg)

		return
	}

	return
}

func Convert(in any, out any) (err error) {
	var (
		tagIn      = "json"
		tagOut     = "sql"
		valueOfIn  = reflect.ValueOf(in)
		valueOfOut = reflect.ValueOf(out)
		typeOfIn   = reflect.TypeOf(in)
		typeOfOut  = reflect.TypeOf(out)
	)

	if valueOfIn.IsNil() || valueOfIn.Kind() != reflect.Pointer {
		return errors.New("In value is nil or not a pointer")
	}

	if valueOfOut.IsNil() || valueOfIn.Kind() != reflect.Pointer {
		return errors.New("Out value is nil or not a pointer")
	}

	for i := 0; i < valueOfIn.Elem().NumField(); i++ {
		for j := 0; j < valueOfOut.Elem().NumField(); j++ {
			if typeOfIn.Elem().Field(i).Tag.Get(tagIn) == typeOfOut.Elem().Field(j).Tag.Get(tagOut) {
				valueOfOut.Elem().Field(i).Set(valueOfIn.Elem().Field(i))
				break
			}
		}
	}

	return
}
