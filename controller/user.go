package controller

import (
	"net/http"
	"praktikum/model"
	"praktikum/service"

	"github.com/labstack/echo/v4"
)

type user struct {
	userSv service.UserSevice
}

// Login implements UserController
func (u *user) Login(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"data":  nil,
			"error": err.Error(),
		})
	}
	userToken, err := u.userSv.Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"data":  nil,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": userToken,
	})
}

// CreateUser implements UserController
func (u *user) CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"data":  nil,
			"error": err.Error(),
		})
	}
	if err := u.userSv.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"data":  nil,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

// GetAllUsers implements UserController
func (u *user) GetAllUsers(c echo.Context) error {
	users, err := u.userSv.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"data":  users,
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": users,
	})
}

func (u *user) GetHallo(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"data": "Ini Hallo Seletah Update",
	})
}

func NewUserController(usv service.UserSevice) UserController {
	return &user{
		userSv: usv,
	}
}
