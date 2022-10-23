package service

import (
	"praktikum/model"
)

type UserSevice interface {
	Login(*model.User) (model.LoginResponse, error)
	CreateUser(*model.User) error
	GetAllUsers() ([]model.User, error)
}
