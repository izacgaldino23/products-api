package product

import (
	"github.com/izacgaldino23/products-api/config/database"
	"github.com/izacgaldino23/products-api/domain"
	"github.com/izacgaldino23/products-api/infra/product"
	"github.com/izacgaldino23/products-api/oops"
	"github.com/izacgaldino23/products-api/utils"
)

func AddProduct(newProduct *Product) (id int64, err error) {
	const msg = "Error on add product"

	tx, err := database.NewTransaction(false)
	if err != nil {
		return id, oops.Wrap(err, msg)
	}

	var (
		productInfra  = product.ProductPS{TX: tx}
		productDomain = domain.Product{}
	)

	if err = utils.Convert(newProduct, &productDomain); err != nil {
		return id, oops.Wrap(err, msg)
	}

	if id, err = productInfra.AddProduct(&productDomain); err != nil {
		return id, oops.Wrap(err, msg)
	}

	if err = tx.Commit(); err != nil {
		return id, oops.Wrap(err, msg)
	}

	return
}

func ListProducts(params *utils.QueryParamList) (out ProductList, err error) {
	const msg = "Error on add product"

	tx, err := database.NewTransaction(true)
	if err != nil {
		return out, oops.Wrap(err, msg)
	}

	var (
		productInfra = product.ProductPS{TX: tx}
		list         domain.ProductList
	)

	if list, err = productInfra.ListProducts(params); err != nil {
		return out, oops.Wrap(err, msg)
	}

	out = ProductList{
		Products: make([]Product, len(list.Products)),
	}

	for i := range list.Products {
		if err = utils.Convert(&list.Products[i], &out.Products[i]); err != nil {
			return out, oops.Wrap(err, msg)
		}
	}

	out.Next, out.Count = list.Next, list.Count

	return
}
