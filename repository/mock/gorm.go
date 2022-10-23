package mock

import (
	"praktikum/model"

	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func (d *DatabaseMock) CreateUser(user *model.User) error {
	args := d.Called()

	return args.Error(0)
}

func (d *DatabaseMock) GetAllUsers() ([]model.User, error) {
	args := d.Called()

	return args.Get(0).([]model.User), args.Error(1)
}

func (d *DatabaseMock) Login(user *model.User) error {
	args := d.Called()

	return args.Error(0)
}
