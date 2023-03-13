package product

import (
	"fmt"

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
		columns []string
	)

	query := c.TX.Builder.
		Select().From(domain.GetTableName(&product))

	isTotal := params.ParseToQuery(&query, map[string]utils.Filter{
		"id":             utils.NewFilter("id", utils.FlagIn),
		"name":           utils.NewFilter("name ilike '%:name%'", utils.FlagEq),
		"created_at_lte": utils.NewFilter("created_at < :created_at::TIMESTAMPTZ", utils.FlagEq),
		"created_at_gte": utils.NewFilter("created_at > :created_at::TIMESTAMPTZ", utils.FlagEq),
	})

	if !isTotal {
		columns, err = utils.GetSqlColumnList(&product)
		if err != nil {
			return out, oops.Err(err)
		}

		query = query.Columns(columns...)
	}

	fmt.Println(query.ToSql())

	rows, err := query.Query()
	if err != nil {
		return out, oops.Err(err)
	}

	out = domain.ProductList{
		Products: []domain.Product{},
	}

	for rows.Next() {
		product = domain.Product{}
		if err = rows.Scan(&product.ID, &product.Name, &product.ImageURL, &product.Link, &product.LastBuyPrice, &product.LastSellPrice, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return out, oops.Err(err)
		}

		out.Products = append(out.Products, product)
	}

	return
}
