package util

import (
	"server/helper"

	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	HashPassword(password string) string 
	ValidatePassword(hashedPassword, password string) error 
}

type HashImpl struct {
	Salt int 
}

func NewHash() Hash {
	return &HashImpl{
		Salt: bcrypt.DefaultCost,
	}
}

func (b *HashImpl) HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	return string(hashPassword)
}


func (b *HashImpl) ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}