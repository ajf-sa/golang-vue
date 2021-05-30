package sgorm

import (
	"github.com/alfuhigi/store-ajf-sa/pkg/user"
	"gorm.io/gorm"
)

type uRepository struct {
	db *gorm.DB
}

func NewGormURepository(db *gorm.DB) user.URepository {
	return &uRepository{
		db,
	}
}
func (r *uRepository) Create(u *user.User) (*user.User, error) {
	return nil, nil
}
func (r *uRepository) FindByID(id uint) (*user.User, error) {
	return nil, nil
}
func (r *uRepository) FindByUserName(u string) (*user.User, error) {
	return nil, nil
}
func (r *uRepository) FindAll() ([]*user.User, error) {
	return nil, nil
}
func (r *uRepository) IsActive(id uint) (bool, error) {
	return false, nil
}
func (r *uRepository) IsStaff(id uint) (bool, error) {
	return false, nil
}
func (r *uRepository) IsEmailVerified(id uint) (bool, error) {
	return false, nil
}
func (r *uRepository) IsPhoneVerified(id uint) (bool, error) {
	return false, nil
}
