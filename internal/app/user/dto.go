package user

import (
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CreateUser struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Gender    string `validate:"oneof=M F prefer_not_to"`
	IPAddress string `validate:"required,ip"`
	Password  string `validate:"required"`
}

type PaginatedResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Data     interface{} `json:"data,omitempty"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

func NewBaseResponse(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Send(c echo.Context, code int, message string, data interface{}) error {
	response := NewBaseResponse(code, message, data)
	return c.JSON(code, response)
}
