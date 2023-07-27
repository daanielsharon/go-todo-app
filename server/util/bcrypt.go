package util

import (
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.Salt)
	helper.PanicIfError(err)

	return string(hashPassword)
}

func (b *BcryptServiceImpl) ValidatePassword(hashedPassword, password string) error {
	hashedPasswordSalted := b.HashPassword(password)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordSalted), []byte(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}