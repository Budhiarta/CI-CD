package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"praktikum/model"
	"praktikum/service/mock"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteUser struct {
	suite.Suite
	controller *user
	mock       *mock.UserMock
}

func (s *suiteUser) SetupSuite() {
	mock := &mock.UserMock{}
	s.mock = mock

	s.controller = &user{
		userSv: mock,
	}
}

func (s *suiteUser) TestGetAllUsers() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		ExpectedBody       []model.User
		ExpectedErr        bool
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			ExpectedBody: []model.User{
				{
					Email:    "rian@gmail.com",
					Password: "rian",
				},
			},
			ExpectedErr: false,
		},
		{
			Name:               "error",
			ExpectedStatusCode: http.StatusInternalServerError,
			ExpectedBody:       []model.User{{}},
			ExpectedErr:        true,
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			mockCall := s.mock.On("GetAllUsers")
			if !v.ExpectedErr {
				mockCall.Return([]model.User{
					{
						Email:    "rian@gmail.com",
						Password: "rian",
					},
				}, nil)
			} else {
				mockCall.Return([]model.User{{}}, errors.New("error"))
			}

			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			w := httptest.NewRecorder()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users")

			err := s.controller.GetAllUsers(ctx)
			s.NoError(err)

			// response body to struct
			rest := struct {
				Data  []model.User
				Error interface{}
			}{}

			err = json.NewDecoder(w.Result().Body).Decode(&rest)

			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)
			s.Equal(v.ExpectedBody, rest.Data)

			mockCall.Unset()
		})
	}
}

func (s *suiteUser) TestGetCreateUser() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Body               map[string]interface{}
		ExpectedErr        bool
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Body: map[string]interface{}{
				"Email":    "test@gmail.com",
				"Password": "test",
			},
			ExpectedErr: false,
		},
		{
			Name:               "Bad Request",
			ExpectedStatusCode: http.StatusBadRequest,
			Body: map[string]interface{}{
				"Email":    1,
				"Password": 2,
			},
			ExpectedErr: true,
		},
		{
			Name:               "Server Error",
			ExpectedStatusCode: http.StatusInternalServerError,
			Body: map[string]interface{}{
				"Email":    "test@gmail.com",
				"Password": "test",
			},
			ExpectedErr: true,
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			mockCall := s.mock.On("CreateUser")
			if !v.ExpectedErr {
				mockCall.Return(nil)
			} else {
				mockCall.Return(errors.New("error"))
			}

			e := echo.New()
			body, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users")

			err := s.controller.CreateUser(ctx)
			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)

			mockCall.Unset()
		})
	}
}
func (s *suiteUser) TestLogin() {
	testCase := []struct {
		Name               string
		ExpectedStatusCode int
		Body               map[string]interface{}
		ExpectedBody       model.LoginResponse
		ExpectedErr        bool
	}{
		{
			Name:               "success",
			ExpectedStatusCode: http.StatusOK,
			Body: map[string]interface{}{
				"Email":    "test@gmail.com",
				"Password": "test",
			},
			ExpectedBody: model.LoginResponse{
				Email: "test@gmail.com",
				Token: "12345",
			},
			ExpectedErr: false,
		},
		{
			Name:               "bad request",
			ExpectedStatusCode: http.StatusBadRequest,
			Body: map[string]interface{}{
				"Email":    1,
				"Password": 1,
			},
			ExpectedBody: model.LoginResponse{},
			ExpectedErr:  true,
		},
		{
			Name:               "record not found",
			ExpectedStatusCode: http.StatusBadRequest,
			Body: map[string]interface{}{
				"Email":    "test@gmail.com",
				"Password": "test1",
			},
			ExpectedBody: model.LoginResponse{},
			ExpectedErr:  true,
		},
	}

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			mockCall := s.mock.On("Login")
			if !v.ExpectedErr {
				mockCall.Return(v.ExpectedBody, nil)
			} else {
				mockCall.Return(model.LoginResponse{}, errors.New("error"))
			}

			e := echo.New()
			body, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/users")

			err := s.controller.Login(ctx)
			s.NoError(err)

			// response body to struct
			rest := struct {
				Data  model.LoginResponse
				Error interface{}
			}{}

			err = json.NewDecoder(w.Result().Body).Decode(&rest)

			s.NoError(err)

			s.Equal(v.ExpectedStatusCode, w.Result().StatusCode)
			s.Equal(v.ExpectedBody, rest.Data)

			mockCall.Unset()
		})
	}
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUser))
}
