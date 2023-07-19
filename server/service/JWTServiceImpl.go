package service

import (
	"fmt"
	"server/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthImpl struct {
	PrivateKey string
}

type JWTClaims struct {
	Username string
	jwt.RegisteredClaims
}

func NewJWTAuthService() JWTService {
	key := helper.GetEnv("SECRET_KEY")
	return &JWTAuthImpl{
		PrivateKey: key,
	}
}

func (s *JWTAuthImpl) TokenGenerate(username string) string {
	claims := &JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	if s.PrivateKey == "" {
		err := fmt.Errorf(".env has not been set!")
		panic(err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.PrivateKey))

	if err != nil {
		panic(err)
	}

	return ss
}

func (s *JWTAuthImpl) TokenValidate(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
		}

		return []byte(s.PrivateKey), nil
	})
}
