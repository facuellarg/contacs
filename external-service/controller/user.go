package controller

import (
	"github.com/labstack/echo/v4"
)

type UserControllerI interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	DeleteUserById(c echo.Context) error
	UpdateUser(c echo.Context) error
}
