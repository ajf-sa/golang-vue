package user

import "github.com/alfuhigi/store-ajf-sa/pkg/store/order"

type User struct {
	Orders []*order.Order
}
