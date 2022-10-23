package controller

import (
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Login(c echo.Context) error
	GetAllUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	GetHallo(c echo.Context) error
}
