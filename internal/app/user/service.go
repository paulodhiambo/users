package user

import (
	"fmt"
	"log/slog"
)

type UserService interface {
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User, id int) (*User, error)
	DeleteUser(userId int) (*User, error)
	GetUserById(userId int) (*User, error)
	GetAllUsers(page int, items int) ([]*User, error)
}

type userService struct {
	Logger         *slog.Logger
	UserRepository UserRepository
}

func (u userService) CreateUser(user *User) (*User, error) {
	dbUser, err := u.UserRepository.CreateUser(user)
	if err != nil {
		u.Logger.Error(err.Error())
		return nil, err
	}
	u.Logger.Info(fmt.Sprintf("User %v created successfully", dbUser))
	return dbUser, nil
}

func (u userService) UpdateUser(user *User, id int) (*User, error) {
	dbUser, err := u.UserRepository.UpdateUser(user, id)
	if err != nil {
		u.Logger.Error(err.Error())
	}
	u.Logger.Info(fmt.Sprintf("User %v updated successfully", dbUser))
	return dbUser, nil
}

func (u userService) DeleteUser(userId int) (*User, error) {
	user, err := u.GetUserById(userId)
	if err != nil {
		u.Logger.Error(err.Error())
	}
	dbUser, err := u.UserRepository.DeleteUser(user)
	if err != nil {
		u.Logger.Error(err.Error())
	}
	u.Logger.Info(fmt.Sprintf("User %v deleted successfully", dbUser))
	return dbUser, nil
}

func (u userService) GetUserById(userId int) (*User, error) {
	dbUser, err := u.UserRepository.GetUserByID(userId)
	if err != nil {
		u.Logger.Error(err.Error())
	}
	u.Logger.Info(fmt.Sprintf("User %v retrieved successfully", dbUser))
	return dbUser, nil
}

func (u userService) GetAllUsers(page int, items int) ([]*User, error) {
	users, err := u.UserRepository.GetAllUsers(page, items)
	if err != nil {
		u.Logger.Error(err.Error())
	}
	u.Logger.Info(fmt.Sprintf("User %v retrieved successfully", users))
	return users, nil
}

func NewUserServiceImpl(userRepository UserRepository, logger *slog.Logger) UserService {
	return &userService{
		Logger:         logger,
		UserRepository: userRepository,
	}
}
