package store

import (
	"github.com/alfuhigi/store-ajf-sa/pkg/store/category"
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Categores  []*category.Category
	Name       string
	Descrption string
	Phone      string
	Email      string
	Address    string
	OwnerID    uint
}
