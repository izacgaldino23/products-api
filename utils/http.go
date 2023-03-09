package utils

import (
	"github.com/gofiber/fiber/v2"
)

type QueryParam []any

type QueryParamList map[string]QueryParam

func (p *QueryParamList) AddParam(key string, value any) *QueryParamList {
	if p.HasKey(key) {
		(*p)[key] = append((*p)[key], value)
	} else {
		(*p)[key] = QueryParam{value}
	}

	return p
}

func (p *QueryParamList) HasKey(key string) bool {
	_, ok := (*p)[key]

	return ok
}

func ParseParams(c *fiber.Ctx) (params QueryParamList) {
	args := c.Context().QueryArgs()
	params = make(QueryParamList)

	args.VisitAll(func(key, value []byte) {
		params.AddParam(string(key), string(value))
	})

	return
}
