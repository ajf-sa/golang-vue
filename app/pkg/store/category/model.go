package category

import "github.com/alfuhigi/store-ajf-sa/pkg/store/product"

type Category struct {
	Products []*product.Product
	Name     string
	Slug     string
	StoreId  uint
}
