package database

import (
	"fmt"
	"praktikum/config"
	"praktikum/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	address := config.Cfg.DB_HOST
	user := config.Cfg.DB_USER
	pass := config.Cfg.DB_PASSWORD
	database := config.Cfg.DB_NAME

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, database)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
