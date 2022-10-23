package lib

import (
	"praktikum/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Library interface {
	GenerateToken(id int) (string, error)
}

type token struct {
}

// GenerateToken implements Library
func (*token) GenerateToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.TOKEN_SECRET))
}

func NewLibrary() Library {
	return &token{}
}
