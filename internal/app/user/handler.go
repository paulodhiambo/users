package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserHandler interface {
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userHandler struct {
	UserService UserService
}

func (u userHandler) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return Send(c, http.StatusBadRequest, "User ID is required", nil)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Send(c, http.StatusBadRequest, "Invalid User ID format", nil)
	}

	user, err := u.UserService.GetUserById(id)
	if err != nil {
		return Send(c, http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve user %v", err), nil)
	}

	return Send(c, http.StatusOK, "User retrieved successfully", user)
}

func (u userHandler) GetUsers(c echo.Context) error {
	pageStr := c.QueryParam("page")
	sizeStr := c.QueryParam("size")
	page, err := strconv.Atoi(pageStr)
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return Send(c, http.StatusBadRequest, "Invalid Page or Size", nil)
	}

	users, err := u.UserService.GetAllUsers(page, size)

	if err != nil {
		return Send(c, http.StatusInternalServerError, fmt.Sprintf("Failed to get users - %v", err), nil)
	}

	pageResponse := Pagination{
		Data:     users,
		Page:     page,
		PageSize: size,
	}

	return Send(c, http.StatusOK, "Users retrieved successfully", pageResponse)
}

func (u userHandler) CreateUser(c echo.Context) error {
	var user CreateUser
	if err := c.Bind(&user); err != nil {
		return Send(c, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err), nil)
	}

	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		return Send(c, http.StatusBadRequest, fmt.Sprintf("Validation failed: %v", err), nil)
	}
	hash, err := HashPassword(user.Password)

	dbUser := &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Gender:    user.Gender,
		Password:  hash,
	}

	createdUser, err := u.UserService.CreateUser(dbUser)
	if err != nil {
		return Send(c, http.StatusInternalServerError, fmt.Sprintf("Failed to create user - %v", err), nil)
	}

	return Send(c, http.StatusCreated, "User created successfully", createdUser)
}

func (u userHandler) UpdateUser(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return Send(c, http.StatusBadRequest, "User ID is required", nil)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Send(c, http.StatusBadRequest, "Invalid User ID format", nil)
	}

	var user CreateUser
	if err := c.Bind(&user); err != nil {
		return Send(c, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err), nil)
	}

	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		return Send(c, http.StatusBadRequest, fmt.Sprintf("Validation failed: %v", err), nil)
	}
	hash, err := HashPassword(user.Password)
	dbUser := &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Gender:    user.Gender,
		Password:  hash,
	}
	updatedUser, err := u.UserService.UpdateUser(dbUser, id)
	if err != nil {
		return Send(c, http.StatusInternalServerError, fmt.Sprintf("Failed to update user - %v", err), nil)
	}
	return Send(c, http.StatusOK, "User updated successfully", updatedUser)
}

func (u userHandler) DeleteUser(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		return Send(c, http.StatusBadRequest, "User ID is required", nil)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Send(c, http.StatusBadRequest, "Invalid User ID format", nil)
	}

	user, err := u.UserService.DeleteUser(id)
	if err != nil {
		return Send(c, http.StatusInternalServerError, fmt.Sprintf("Failed to delete user - %v", err), nil)
	}

	return Send(c, http.StatusOK, "User deleted successfully", user)
}

func NewUserHandler(userService UserService) UserHandler {
	return &userHandler{
		UserService: userService,
	}
}
