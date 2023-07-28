package util

import (
	"fmt"
	"server/helper"

	"golang.org/x/crypto/bcrypt"
)

type BcryptService interface {
	HashPassword(password string) string 
	ValidatePassword(hashedPassword, password string) error 
}

type BcryptServiceImpl struct {
	Salt int 
}

func NewBcrypt() BcryptService {
	return &BcryptServiceImpl{
		Salt: bcrypt.DefaultCost,
	}
}

func (b *BcryptServiceImpl) HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	return string(hashPassword)
}

func (b *BcryptServiceImpl) ValidatePassword(hashedPassword, password string) error {
	hashedPasswordSalted := b.HashPassword(password)
	fmt.Println("database password", hashedPassword)
	fmt.Println("user request password hashed", hashedPasswordSalted)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}