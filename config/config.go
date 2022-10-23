package config

import "os"

type config struct {
	API_PORT     string
	TOKEN_SECRET string
	DB_HOST      string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
}

var Cfg *config

func InitConfig() {
	cfg := &config{}

	cfg.API_PORT = SetEnv("API_PORT", ":8000")
	cfg.TOKEN_SECRET = SetEnv("TOKEN_SECRET", "ADsf78saAA")
	cfg.DB_HOST = SetEnv("DB_HOST", "mysql-mysql-1:3306")
	cfg.DB_USER = SetEnv("DB_USER", "root")
	cfg.DB_PASSWORD = SetEnv("DB_PASSWORD", "root123")
	cfg.DB_NAME = SetEnv("DB_NAME", "praktikum")

	Cfg = cfg
}

func SetEnv(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}

	return val
}
