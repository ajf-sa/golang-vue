package user

type URepository interface {
	Create(u *User) (*User, error)
	FindByID(id uint) (*User, error)
	FindByUserName(u string) (*User, error)
	FindAll() ([]*User, error)

	IsActive(id uint) (bool, error)
	IsStaff(id uint) (bool, error)
	IsEmailVerified(id uint) (bool, error)
	IsPhoneVerified(id uint) (bool, error)
}
