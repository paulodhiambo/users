package user

import "log/slog"

type UserService interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(user *User) (*User, error)
	GetUserById(userId int) (*User, error)
	GetAllUsers() ([]*User, error)
}

type userService struct {
	Logger         *slog.Logger
	UserRepository UserRepository
}

func (u userService) CreateUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) UpdateUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) DeleteUser(user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetUserById(userId int) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetAllUsers() ([]*User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserServiceImpl(userRepository UserRepository, logger *slog.Logger) UserService {
	return &userService{
		Logger:         logger,
		UserRepository: userRepository,
	}
}
