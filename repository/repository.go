package repository

import (
	"praktikum/model"
)

type Database interface {
	Login(*model.User) error
	CreateUser(*model.User) error
	GetAllUsers() ([]model.User, error)
}
