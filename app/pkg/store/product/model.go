package product

import (
	"github.com/alfuhigi/store-ajf-sa/pkg/store/product/item"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID  uint
	Name        string
	Slug        string
	Sku         string
	Description string
	Content     string
	Items       []*item.Item
	Price       float32
	Qty         int
	IsActive    bool
	IsAvilable  bool
}
