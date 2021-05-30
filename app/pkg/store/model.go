package store

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name       string
	Descrption string
	Phone      string
	Email      string
	Address    string
	OwnerID    uint
}
