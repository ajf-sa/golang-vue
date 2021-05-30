package user

type UService interface {
	CreateUser(u *User) (*User, error)
	FindUserByID(u *User) (*User, error)
	FindUserByUserName(username string) (*User, error)
	FindAllUser() ([]*User, error)
	IsActiveUser(id uint) (bool, error)
	IsStaffUser(id uint) (bool, error)
	IsEmailVerifiedUser(id uint) (bool, error)
	IsPhoneVerifiedUser(id uint) (bool, error)
}

type uService struct {
	repo URepository
}

func NewUService(repo URepository) UService {
	return &uService{
		repo,
	}
}

func (s *uService) CreateUser(u *User) (*User, error) {
	return nil, nil
}

func (s *uService) FindUserByID(u *User) (*User, error) {
	return nil, nil
}

func (s *uService) FindUserByUserName(username string) (*User, error) {
	return nil, nil
}

func (s *uService) FindAllUser() ([]*User, error) {
	return nil, nil
}

func (s *uService) IsActiveUser(id uint) (bool, error) {
	return false, nil
}

func (s *uService) IsStaffUser(id uint) (bool, error) {
	return false, nil
}

func (s *uService) IsEmailVerifiedUser(id uint) (bool, error) {
	return false, nil
}

func (s *uService) IsPhoneVerifiedUser(id uint) (bool, error) {
	return false, nil
}
