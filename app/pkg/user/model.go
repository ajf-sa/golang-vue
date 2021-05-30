package user

import (
	"github.com/alfuhigi/store-ajf-sa/pkg/store"
	"github.com/alfuhigi/store-ajf-sa/pkg/store/order"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Orders          []*order.Order
	Stores          []*store.Store
	Username        string
	Password        []byte
	Email           string
	Phone           string
	IsStaff         bool
	IsActive        bool
	IsPhoneVerified bool
	IsEmailVerified bool
}
