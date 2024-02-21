package product

import (
	"github.com/izacgaldino23/products-api/config/database"
	"github.com/izacgaldino23/products-api/domain"
	"github.com/izacgaldino23/products-api/infra/default_db"
	"github.com/izacgaldino23/products-api/oops"
	"github.com/izacgaldino23/products-api/utils"
)

type ProductPS struct {
	TX    *database.Transaction
	Model domain.Product
}

func (c *ProductPS) AddProduct(product *domain.Product) (id int64, err error) {
	return default_db.Insert(product, c.TX)
}

func (c *ProductPS) UpdateProduct(product *domain.Product) (err error) {
	return default_db.Update(product, *product.ID, c.TX)
}

func (c *ProductPS) DeleteProduct(id int64) (err error) {
	return default_db.Delete(&domain.Product{}, id, c.TX)
}

func (c *ProductPS) ListProducts(params *utils.QueryParamList) (out domain.ProductList, err error) {
	var (
		product = domain.Product{}
		total   int64
		next    bool
		result  []interface{}
	)

	if !params.HasKey("removed") {
		params.AddParam("removed", false)
	}

	query := c.TX.Builder.
		Select().From(domain.GetTableName(&product))

	if total, next, result, err = params.MakeQuery(&query, map[string]utils.Filter{
		"id":             utils.NewFilter("id", utils.FlagIn),
		"name":           utils.NewFilter("name ilike '%:name%'", utils.FlagEq),
		"created_at_lte": utils.NewFilter("created_at < :created_at::TIMESTAMPTZ", utils.FlagEq),
		"created_at_gte": utils.NewFilter("created_at > :created_at::TIMESTAMPTZ", utils.FlagEq),
		"removed":        utils.NewFilter("removed_at", utils.FlagNil),
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

func (c *ProductPS) GetProductByField(field string, value any) (*domain.Product, error) {
	return default_db.SelectByField(&c.Model, field, value, c.TX)
}
