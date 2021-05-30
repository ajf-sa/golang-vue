package item

import (
	"image"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ProductID  uint
	Name       string
	Slug       string
	Sku        string
	Color      string
	Images     []*image.Image
	Size       string
	Qty        int
	IsActive   bool
	IsAvilable bool
}
