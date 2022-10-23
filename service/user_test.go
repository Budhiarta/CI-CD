package service

import (
	"errors"
	mockL "praktikum/lib/mock"
	"praktikum/model"
	"praktikum/repository/mock"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suiteService struct {
	suite.Suite
	service *user
	mock    *mock.DatabaseMock
	mockLib *mockL.LibMock
}

func (s *suiteService) SetupSuite() {
	mock := &mock.DatabaseMock{}
	mockL := &mockL.LibMock{}
	s.mock = mock
	s.mockLib = mockL

	s.service = &user{
		repo: mock,
		lib:  mockL,
	}
}

func (s *suiteService) TestCreateUser() {
	testCase := []struct {
		Name        string
		Body        model.User
		ExpectedErr error
	}{
		{
			Name: "success",
			Body: model.User{
				Email:    "test@gmail.com",
				Password: "test",
			},
			ExpectedErr: nil,
		},
		{
			Name: "error",
			Body: model.User{
				Email:    "test@gmail.com",
				Password: "test",
			},
			ExpectedErr: errors.New("error"),
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			data := v.Body
			mockCall := s.mock.On("CreateUser")
			if v.ExpectedErr != nil {
				mockCall.Return(errors.New("error"))
			} else {
				mockCall.Return(nil)
			}

			err := s.service.CreateUser(&data)

			s.Equal(v.ExpectedErr, err)

			mockCall.Unset()
		})
	}
}

func (s *suiteService) TestGetAllUsers() {
	testCase := []struct {
		Name           string
		ExpectedResult []model.User
		ExpectedErr    error
	}{
		{
			Name: "success",
			ExpectedResult: []model.User{
				{
					Email:    "test@gmail.com",
					Password: "test",
				},
			},
			ExpectedErr: nil,
		},
		{
			Name:           "error",
			ExpectedResult: []model.User{{}},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			mockCall := s.mock.On("GetAllUsers").Return(v.ExpectedResult, v.ExpectedErr)
			res, err := s.service.GetAllUsers()

			s.Equal(v.ExpectedResult, res)
			s.Equal(v.ExpectedErr, err)

			mockCall.Unset()
		})
	}
}
func (s *suiteService) TestLogin() {
	testCase := []struct {
		Name           string
		Body           model.User
		ExpectedResult model.LoginResponse
		ExpectedErr    error
	}{
		{
			Name: "success",
			Body: model.User{
				Email:    "test@gmail.com",
				Password: "test",
			},
			ExpectedResult: model.LoginResponse{
				Email: "test@gmail.com",
				Token: "1234",
			},
			ExpectedErr: nil,
		},
		{
			Name: "error database",
			Body: model.User{
				Email:    "test@gmail.com",
				Password: "test",
			},
			ExpectedResult: model.LoginResponse{},
			ExpectedErr:    errors.New("error"),
		},
		{
			Name: "token error",
			Body: model.User{
				Email:    "test@gmail.com",
				Password: "test",
			},
			ExpectedResult: model.LoginResponse{},
			ExpectedErr:    errors.New("error"),
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			mockCall := s.mock.On("Login").Return(v.ExpectedErr)
			mockCall2 := s.mockLib.On("GenerateToken").Return("1234", v.ExpectedErr)

			res, err := s.service.Login(&v.Body)
			s.Equal(v.ExpectedResult, res)
			s.Equal(v.ExpectedErr, err)

			mockCall.Unset()
			mockCall2.Unset()
		})
	}
}

func TestSuiteServices(t *testing.T) {
	suite.Run(t, new(suiteService))
}
