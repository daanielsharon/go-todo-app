package util

import (
	"server/helper"

	"golang.org/x/crypto/bcrypt"
)

type HashCrypt interface {
	HashPassword(password string) string 
	ValidatePassword(hashedPassword, password string) error 
}

type HashcryptImpl struct {
	Salt int 
}

func NewHashcrypt() HashCrypt {
	return &HashcryptImpl{
		Salt: bcrypt.DefaultCost,
	}
}

func (b *HashcryptImpl) HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	return string(hashPassword)
}


func (b *HashcryptImpl) ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}