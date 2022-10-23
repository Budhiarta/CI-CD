package repository

import (
	"praktikum/model"

	"gorm.io/gorm"
)

type gormSql struct {
	db *gorm.DB
}

// CreateUser implements Database
func (g *gormSql) CreateUser(user *model.User) error {
	err := g.db.Create(&user).Error
	return err
}

// GetAllUsers implements Database
func (g *gormSql) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := g.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

// Login implements Database
func (g *gormSql) Login(user *model.User) error {
	err := g.db.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error

	return err
}

func NewGromDB(db *gorm.DB) Database {
	return &gormSql{
		db: db,
	}
}
