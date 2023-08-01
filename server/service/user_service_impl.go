package service

import (
	"context"
	"database/sql"
	"server/exception"
	"server/helper"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	containerTodo "server/repository/todo/container"
	"server/util"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	ItemRepository containerTodo.ContainerRepository
	DB             *sql.DB
	Timeout        time.Duration
	Validate       *validator.Validate
}

func NewUserService(u repository.UserRepository, t containerTodo.ContainerRepository, db *sql.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: u,
		ItemRepository: t,
		DB:             db,
		Timeout:        time.Duration(2) * time.Second,
		Validate:       validator,
	}
}

func (s *UserServiceImpl) Create(c context.Context, req *web.UserCreateRequest) *web.UserCreateResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	hashedPassword := util.NewHash().HashPassword(req.Password)
	user := domain.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	u, _ := s.UserRepository.Find(ctx, s.DB, &user)

	if u.Username != "" {
		panic(exception.NewValidationError("Duplicate username"))
	}

	r, err := s.UserRepository.Save(ctx, s.DB, &user)
	helper.PanicIfError(err)

	err = s.ItemRepository.InitGroup(ctx, s.DB, int(r.ID))
	helper.PanicIfError(err)

	response := web.UserCreateResponse{
		ID:       r.ID,
		Username: r.Username,
	}

	return &response
}

func (s *UserServiceImpl) Get(c context.Context, req *web.UserGetRequest) *web.UserGetResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
		Password: req.Password,
	}

	u, err := s.UserRepository.Find(ctx, s.DB, &user)
	if err != nil {
		panic(exception.NewValidationError(err.Error()))
	}

	err = util.NewHash().ValidatePassword(u.Password, req.Password)
	if err != nil {
		panic(exception.NewNotFoundError("Username or password is wrong"))
	}

	response := &web.UserGetResponse{
		ID:       u.ID,
		Username: u.Username,
	}

	return response
}
