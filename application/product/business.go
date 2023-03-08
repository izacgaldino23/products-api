package product

import (
	"fmt"
)

func AddProduct(product *Product) (id int, err error) {
	fmt.Println(*product)

	return
}
