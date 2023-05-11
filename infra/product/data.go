package product

import (
	"github.com/izacgaldino23/products-api/config/database"
	"github.com/izacgaldino23/products-api/domain"
	"github.com/izacgaldino23/products-api/oops"
	"github.com/izacgaldino23/products-api/utils"
)

type ProductPS struct {
	TX *database.Transaction
}

func (c *ProductPS) AddProduct(product *domain.Product) (id int64, err error) {
	valueMap, err := utils.ParseStructToMap(product)
	if err != nil {
		return
	}

	if err = c.TX.Builder.
		Insert(domain.GetTableName(product)).
		SetMap(valueMap).
		Suffix("RETURNING id").
		Scan(&id); err != nil {
		return id, oops.Err(err)
	}

	return
}

func (c *ProductPS) ListProducts(params *utils.QueryParamList) (out domain.ProductList, err error) {
	var (
		product = domain.Product{}
		total   int64
		next    bool
		result  []interface{}
	)

	query := c.TX.Builder.
		Select().From(domain.GetTableName(&product))

	if total, next, result, err = params.MakeQuery(&query, map[string]utils.Filter{
		"id":             utils.NewFilter("id", utils.FlagIn),
		"name":           utils.NewFilter("name ilike '%:name%'", utils.FlagEq),
		"created_at_lte": utils.NewFilter("created_at < :created_at::TIMESTAMPTZ", utils.FlagEq),
		"created_at_gte": utils.NewFilter("created_at > :created_at::TIMESTAMPTZ", utils.FlagEq),
	}, &product); err != nil {
		return out, oops.Err(err)
	}

	out = domain.ProductList{
		Products: []domain.Product{},
		Count:    total,
		Next:     next,
	}

	for i := range result {
		out.Products = append(out.Products, *result[i].(*domain.Product))
	}

	return
}
