package service

import (
	"context"
	"database/sql"
	"server/exception"
	"server/helper"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	TodoRepository repository.TodoRepository
	DB             *sql.DB
	Timeout        time.Duration
	Validate       *validator.Validate
}

func NewUserService(u repository.UserRepository, t repository.TodoRepository, db *sql.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: u,
		TodoRepository: t,
		DB:             db,
		Timeout:        time.Duration(2) * time.Second,
		Validate:       validator,
	}
}

func (s *UserServiceImpl) CreateUsername(c context.Context, req *web.UserCreateUsernameRequest) *web.UserCreateUsernameResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	u, _ := s.UserRepository.FindUsername(ctx, s.DB, &user)

	if u.Username != "" {
		panic(exception.NewValidationError("Duplicate username"))
	}

	r, err := s.UserRepository.SaveUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	err = s.TodoRepository.InitGroup(ctx, s.DB, int(r.ID))
	helper.PanicIfError(err)

	response := web.UserCreateUsernameResponse{
		ID:       r.ID,
		Username: r.Username,
	}

	return &response
}

func (s *UserServiceImpl) GetUsername(c context.Context, req *web.UserGetUsernameRequest) *web.UserGetUsernameResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	u, err := s.UserRepository.FindUsername(ctx, s.DB, &user)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	response := &web.UserGetUsernameResponse{
		ID:       u.ID,
		Username: u.Username,
	}

	return response
}
