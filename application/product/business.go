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

	if err = utils.Convert(newProduct, &productDomain, false); err != nil {
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
		if err = utils.Convert(&list.Products[i], &out.Products[i], true); err != nil {
			return out, oops.Wrap(err, msg)
		}
	}

	out.Next, out.Count = list.Next, list.Count

	return
}

func GetProduct(id int64) (out Product, err error) {
	const msg = "Error on get product"

	tx, err := database.NewTransaction(true)
	if err != nil {
		return out, oops.Wrap(err, msg)
	}

	var (
		productInfra  = product.ProductPS{TX: tx}
		productFromDB *domain.Product
	)

	if productFromDB, err = productInfra.GetProductByField("id", id); err != nil {
		return out, oops.Wrap(err, msg)
	}

	if err = utils.Convert(productFromDB, &out, true); err != nil {
		return out, oops.Wrap(err, msg)
	}

	return
}

func UpdateProduct(id int64, productUpdate *Product) (err error) {
	const msg = "Error on update product"

	tx, err := database.NewTransaction(false)
	if err != nil {
		return oops.Wrap(err, msg)
	}

	var (
		productInfra  = product.ProductPS{TX: tx}
		productDomain = domain.Product{}
	)

	if err = utils.Convert(productUpdate, &productDomain, false); err != nil {
		return oops.Wrap(err, msg)
	}

	productDomain.ID = &id
	if err = productInfra.UpdateProduct(&productDomain); err != nil {
		return oops.Wrap(err, msg)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msg)
	}

	return
}

func DeleteProduct(id int64) (err error) {
	const msg = "Error on delete product"

	tx, err := database.NewTransaction(false)
	if err != nil {
		return oops.Wrap(err, msg)
	}

	var (
		productInfra = product.ProductPS{TX: tx}
	)

	if err = productInfra.DeleteProduct(id); err != nil {
		return oops.Wrap(err, msg)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msg)
	}

	return
}
