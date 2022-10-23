package main

import (
	"praktikum/config"
	"praktikum/database"
	"praktikum/routes"
)

func main() {
	config.InitConfig()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	if err := database.MigrateDB(db); err != nil {
		panic(err)
	}

	e := routes.InitRoutes(db)

	e.Logger.Fatal(e.Start(config.Cfg.API_PORT))
}
