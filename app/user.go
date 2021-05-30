package main

import (
	"github.com/alfuhigi/store-ajf-sa/database/sgorm"
	"github.com/alfuhigi/store-ajf-sa/pkg/user"
	"gorm.io/gorm"
)

func userH(db *gorm.DB) user.UHandler {
	repo := sgorm.NewGormURepository(db)
	service := user.NewUService(repo)
	handler := user.NewUHandler(service)
	return handler
}
