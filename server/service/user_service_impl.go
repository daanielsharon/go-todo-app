package service

import (
	"context"
	"database/sql"
	"errors"
	"server/helper"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Timeout        time.Duration
}

func NewUserService(u repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: u,
		DB:             db,
		Timeout:        time.Duration(2) * time.Second,
	}
}

func (s *UserServiceImpl) CreateUsername(c context.Context, req *web.UserCreateUsernameRequest) *web.UserCreateUsernameResponse {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	u, err := s.UserRepository.FindUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	if u != nil {
		panic(errors.New("duplicate username"))
	}

	r, err := s.UserRepository.SaveUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	response := web.UserCreateUsernameResponse{
		ID:       r.ID,
		Username: r.Username,
	}

	return &response
}

func (s *UserServiceImpl) GetUsername(c context.Context, req *web.UserGetUsernameRequest) *web.UserGetUsernameResponse {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	u, err := s.UserRepository.FindUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	response := &web.UserGetUsernameResponse{
		ID:       u.ID,
		Username: u.Username,
	}

	return response
}
