package details

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ProductID uint
	OrderID   uint
	Qty       int
	UnitPrice float32
}
