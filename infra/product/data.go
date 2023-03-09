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
