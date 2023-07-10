package service

import (
	"context"
	"database/sql"
	"errors"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	db             *sql.DB
	timeout        time.Duration
}

func NewUserService(u repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: u,
		db:             db,
		timeout:        time.Duration(2) * time.Second,
	}
}

func (s *UserServiceImpl) CreateUsername(c context.Context, req *web.UserCreateUsernameRequest) (*web.UserCreateUsernameResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	u, err := s.UserRepository.FindUsername(ctx, s.db, &user)
	if err != nil {
		return nil, err
	}

	if u != nil {
		return nil, errors.New("duplicate username")
	}

	r, err := s.UserRepository.SaveUsername(ctx, s.db, &user)

	response := web.UserCreateUsernameResponse{
		ID:       r.ID,
		Username: r.Username,
	}

	return &response, nil
}

func (s *UserServiceImpl) GetUsername(c context.Context, req *web.UserGetUsernameRequest) (*web.UserGetUsernameResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	u, err := s.UserRepository.FindUsername(ctx, s.db, &user)
	if err != nil {
		return &web.UserGetUsernameResponse{}, err
	}

	response := &web.UserGetUsernameResponse{
		ID:       u.ID,
		Username: u.Username,
	}

	return response, nil
}
