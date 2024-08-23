package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(*User) (*User, error)
	UpdateUser(user *User, id int) (*User, error)
	DeleteUser(*User) (*User, error)
	GetUserByID(id int) (*User, error)
	GetAllUsers(page int, results int) ([]*User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func (u userRepository) CreateUser(user *User) (*User, error) {
	if err := u.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepository) UpdateUser(user *User, id int) (*User, error) {
	var existingUser User
	if err := u.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.Gender = user.Gender
	existingUser.IPAddress = user.IPAddress
	existingUser.Password = user.Password

	if err := u.DB.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (u userRepository) DeleteUser(user *User) (*User, error) {
	var existingUser User
	if err := u.DB.First(&existingUser, user.ID).Error; err != nil {
		return nil, err
	}
	if err := u.DB.Delete(&existingUser).Error; err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func (u userRepository) GetUserByID(id int) (*User, error) {
	var existingUser User
	if err := u.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}
	return &existingUser, nil
}

func (u userRepository) GetAllUsers(page int, results int) ([]*User, error) {
	var users []*User
	offset := (page - 1) * results
	if err := u.DB.Limit(results).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}
