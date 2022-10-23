package mock

import (
	"praktikum/model"

	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (u *UserMock) GetAllUsers() ([]model.User, error) {
	args := u.Called()

	return args.Get(0).([]model.User), args.Error(1)
}

func (u *UserMock) CreateUser(*model.User) error {
	args := u.Called()

	return args.Error(0)
}

func (u *UserMock) Login(*model.User) (model.LoginResponse, error) {
	args := u.Called()

	return args.Get(0).(model.LoginResponse), args.Error(1)
}
