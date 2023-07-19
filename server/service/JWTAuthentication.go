package service

import "github.com/golang-jwt/jwt/v5"

type JWTService interface {
	TokenGenerate(username string) string
	TokenValidate(token string) (*jwt.Token, error)
}
