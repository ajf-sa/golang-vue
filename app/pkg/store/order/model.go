package order

import (
	"github.com/alfuhigi/store-ajf-sa/pkg/store/order/details"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint
	StoreID    uint
	Status     []*Status
	Items      []*details.Item
	Address    string
	Phone      string
	Email      string
	TotalPrice float32
}

type Status struct {
	gorm.Model
	Name    string
	OrderID uint
	Content string
}
