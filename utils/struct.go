package utils

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

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

func GetFieldList(m interface{}) (fields []any, err error) {
	var (
		typeOf  = reflect.TypeOf(m)
		valueOf = reflect.ValueOf(m)
		tagName = "sql"
	)

	for i := 0; i < typeOf.Elem().NumField(); i++ {
		var (
			fieldName = typeOf.Elem().Field(i).Tag.Get(tagName)
			newField  reflect.Value
			value     any
		)

		if fieldName == "" {
			err = errTagIsEmpty
			return
		} else if fieldName == "-" {
			continue
		}
		fieldType := valueOf.Elem().Field(i).Type()

		switch valueOf.Elem().Field(i).Interface().(type) {
		case time.Time:
			newField = reflect.New(reflect.TypeOf(time.Now()))
		default:
			newField = reflect.New(fieldType)
		}

		value = newField.Elem().Interface()
		fields = append(fields, &value)
	}

	return
}

func FieldsToStruct(fields []any, out interface{}) (newObject interface{}, err error) {
	var (
		tag        = "sql"
		valueOfOut = reflect.ValueOf(out)
		typeOfOut  = reflect.TypeOf(out)
		newThing   = reflect.New(typeOfOut.Elem())
	)

	if valueOfOut.IsNil() || valueOfOut.Kind() != reflect.Pointer {
		return newObject, errors.New("Out value is nil or not a pointer")
	}

	for i := 0; i < newThing.Elem().NumField(); i++ {
		if typeOfOut.Elem().Field(i).Tag.Get(tag) != "" {
			temp := fields[i].(*interface{})
			v := *temp
			name := typeOfOut.Elem().Field(i).Name

			switch value := v.(type) {
			case time.Time:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			case int64:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			case int32:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			case float32:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			case float64:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			case string:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			case bool:
				newThing.Elem().FieldByName(name).Set(reflect.ValueOf(&value))
			default:
				// Nil value
			}
		}
	}

	newObject = newThing.Interface()

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

	if valueOfOut.IsNil() || valueOfOut.Kind() != reflect.Pointer {
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
