package service

import (
	"praktikum/lib"
	"praktikum/model"
	"praktikum/repository"
)

type user struct {
	repo repository.Database
	lib  lib.Library
}

// CreateUser implements UserSevice
func (u *user) CreateUser(data *model.User) error {
	err := u.repo.CreateUser(data)

	return err
}

// GetAllUsers implements UserSevice
func (u *user) GetAllUsers() ([]model.User, error) {
	res, err := u.repo.GetAllUsers()
	if err != nil {
		return res, err
	}

	return res, err
}

// Login implements UserSevice
func (u *user) Login(user *model.User) (model.LoginResponse, error) {
	var res model.LoginResponse
	if err := u.repo.Login(user); err != nil {
		return res, err
	}

	token, err := u.lib.GenerateToken(int(user.ID))
	if err != nil {
		return res, err
	}

	res = model.LoginResponse{
		Email: user.Email,
		Token: token,
	}

	return res, nil
}

func NewUserService(repository repository.Database, library lib.Library) UserSevice {
	return &user{
		repo: repository,
		lib:  library,
	}
}
