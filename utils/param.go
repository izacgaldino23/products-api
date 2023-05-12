package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
)

const (
	FlagEq = iota + 1
	FlagNotEq
	FlagIn
	FlagNotIn
	FlagNil
)

const maxDefaultLimit int64 = 15

type QueryParam []any

type QueryParamList map[string]QueryParam

type Filter struct {
	Condition string
	Flag      int
}

func (q *QueryParam) GetInt(index int) (value int64) {
	temp := fmt.Sprintf("%v", (*q)[index])
	value, _ = strconv.ParseInt(temp, 10, 64)

	return
}

func (q *QueryParam) GetBol() (value bool) {
	temp := fmt.Sprintf("%v", (*q)[0])
	value, _ = strconv.ParseBool(temp)

	return
}

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

func (p *QueryParamList) MakeQuery(query *squirrel.SelectBuilder, filters map[string]Filter, object interface{}) (total int64, next bool, out []interface{}, err error) {
	var (
		totalTag  = "total"
		sizeTag   = "size"
		offsetTag = "offset"
		limit     = maxDefaultLimit
		columns   []string
		fields    []any
		isTotal   bool
	)

	query.PlaceholderFormat(squirrel.Question)

	for i, v := range *p {
		switch i {
		case totalTag:
			*query = query.Column("count(*) as total")
			isTotal = true
		case sizeTag:
			limit = v.GetInt(0)
		case offsetTag:
			query.Offset(uint64(v.GetInt(0)))
		default:
			for j := range filters {
				if j == i {
					condition := strings.Replace(filters[j].Condition, ":"+i, "?", 1)

					switch filters[j].Flag {
					case FlagEq:
						*query = query.Where(squirrel.Eq{
							condition: v[0],
						})
					case FlagNotEq:
						*query = query.Where(squirrel.NotEq{
							condition: v[0],
						})
					case FlagIn:
						*query = query.Where(squirrel.Eq{
							condition: v,
						})
					case FlagNotIn:
						*query = query.Where(squirrel.NotEq{
							condition: v,
						})
					case FlagNil:
						if !v.GetBol() {
							*query = query.Where(squirrel.Eq{
								condition: nil,
							})
						} else {
							*query = query.Where(squirrel.NotEq{
								condition: nil,
							})
						}
					}

					break
				}
			}
		}
	}

	query.
		PlaceholderFormat(squirrel.Dollar)

	if !isTotal {
		columns, err = GetSqlColumnList(object)
		if err != nil {
			return
		}

		*query = query.Columns(columns...)
	} else {
		query.Limit(uint64(limit + 1))

		if err = query.QueryRow().Scan(&total); err != nil {
			return
		}

		return
	}

	result, err := query.Query()
	if err != nil {
		return
	}

	out = make([]interface{}, 0)
	var qtd int64
	for result.Next() {
		var (
			newObject interface{}
		)
		if qtd == limit {
			next = true
			break
		}

		fields, err = GetFieldList(object)
		if err != nil {
			return
		}

		if err = result.Scan(fields...); err != nil {
			return
		}

		if newObject, err = FieldsToStruct(fields, object); err != nil {
			return
		}

		out = append(out, newObject)
		qtd++
	}

	return
}

func ParseParams(c *fiber.Ctx) (params QueryParamList) {
	args := c.Context().QueryArgs()
	params = make(QueryParamList)

	args.VisitAll(func(key, value []byte) {
		params.AddParam(string(key), string(value))
	})

	return
}

func NewFilter(condition string, flag int) (filter Filter) {
	return Filter{condition, flag}
}
