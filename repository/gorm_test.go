package repository

import (
	"errors"
	"praktikum/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type suiteDatabase struct {
	suite.Suite
	mock       sqlmock.Sqlmock
	repository *gormSql
}

func (s *suiteDatabase) SetupSuite() {
	db, mocking, _ := sqlmock.New()

	dbGorm, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	s.mock = mocking
	s.repository = &gormSql{
		db: dbGorm,
	}
}

func (s *suiteDatabase) TestCreateUser() {
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
			if v.ExpectedErr != nil {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)")).WillReturnError(v.ExpectedErr)
				s.mock.ExpectRollback()
			} else {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`password`) VALUES (?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 0))
				s.mock.ExpectCommit()
			}
			data := v.Body
			err := s.repository.CreateUser(&data)

			s.Equal(v.ExpectedErr, err)

		})
	}
}
func (s *suiteDatabase) TestGetAllUsers() {
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
			ExpectedResult: []model.User(nil),
			ExpectedErr:    errors.New("error"),
		},
	}

	row := sqlmock.NewRows([]string{"email", "password"}).AddRow("test@gmail.com", "test")

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			if v.ExpectedErr != nil {
				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).WillReturnError(v.ExpectedErr)
			} else {
				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).WillReturnRows(row)
			}
			res, err := s.repository.GetAllUsers()

			s.Equal(v.ExpectedResult, res)
			s.Equal(v.ExpectedErr, err)

		})
	}
}
func (s *suiteDatabase) TestLogin() {
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
	}
	row := sqlmock.NewRows([]string{"email", "password"}).AddRow("test@gmail.com", "test")

	for _, v := range testCase {
		s.T().Run(v.Name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (email = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).WillReturnRows(row)
			body := v.Body

			err := s.repository.Login(&body)

			s.Equal(v.ExpectedErr, err)

		})
	}
}

func (s *suiteDatabase) TearDownSuite() {
	s.repository = nil
	s.mock = nil
}

func TestSuiteDatabase(t *testing.T) {
	suite.Run(t, new(suiteDatabase))
}
