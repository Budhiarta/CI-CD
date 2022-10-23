package routes

import (
	"praktikum/config"
	"praktikum/controller"
	"praktikum/lib"
	"praktikum/repository"
	"praktikum/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	lib := lib.NewLibrary()
	repo := repository.NewGromDB(db)
	usv := service.NewUserService(repo, lib)
	userC := controller.NewUserController(usv)

	e.POST("/users", userC.CreateUser)
	e.POST("/login", userC.Login)
	e.GET("/hallo", userC.GetHallo)

	m := e.Group("")
	m.Use(middleware.JWT([]byte(config.Cfg.TOKEN_SECRET)))
	m.GET("/users", userC.GetAllUsers)

	return e
}
