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

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	UserRepository repository.UserRepository
	DB             *sql.DB
	Timeout        time.Duration
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, userRepository repository.UserRepository, db *sql.DB, validator *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		UserRepository: userRepository,
		DB:             db,
		Timeout:        time.Duration(2) * time.Second,
		Validate:       validator,
	}
}

func (s *TodoServiceImpl) Create(c context.Context, req *web.TodoCreateRequest) *web.TodoCreateUpdateResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	newUser := &domain.TodoListInsertUpdate{
		Name:    req.Name,
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}

	r, err := s.TodoRepository.Save(ctx, s.DB, newUser)
	if err != nil {
		panic(err)
	}

	res := &web.TodoCreateUpdateResponse{
		ID:      r.ID,
		Name:    r.Name,
		UserID:  r.UserID,
		GroupID: r.GroupID,
	}

	return res
}

func (s *TodoServiceImpl) GetByUsername(c context.Context, req *web.TodoGetRequest) *[]web.TodoGetResponse {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	_, err := s.UserRepository.Find(ctx, s.DB, &user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	res, err := s.TodoRepository.FindByUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	var response []web.TodoGetResponse

	for _, val := range *res {
		var item interface{}

		// if first data is nil, it means there's no data for the entire todo group
		if val.Item[0] == nil {
			item = make([]interface{}, 0)
		} else {
			item = val.Item
		}

		response = append(response, web.TodoGetResponse{
			ID:        val.ID,
			GroupName: val.Name,
			Item:      item,
			Priority:  val.Priority,
		})
	}

	return &response
}

func (s *TodoServiceImpl) Update(c context.Context, req *web.TodoUpdateRequest) *web.TodoCreateUpdateResponse {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	_, err = s.TodoRepository.FindById(ctx, s.DB, req.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	groupReq := &domain.TodoGroup{
		ID:     int64(req.GroupID),
		UserID: req.UserID,
	}

	_, err = s.TodoRepository.FindGroup(ctx, s.DB, groupReq)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	data := &domain.TodoListInsertUpdate{
		ID:      req.ID,
		Name:    req.Name,
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}

	r := s.TodoRepository.Update(ctx, s.DB, data)

	res := &web.TodoCreateUpdateResponse{
		ID:      r.ID,
		Name:    r.Name,
		GroupID: req.GroupID,
		UserID:  r.UserID,
	}

	return res
}

func (s *TodoServiceImpl) Remove(c context.Context, req *web.TodoDeleteRequest) {
	err := s.Validate.Struct(req)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	_, err = s.TodoRepository.FindById(ctx, s.DB, req.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.TodoRepository.Delete(ctx, s.DB, &domain.TodoList{ID: req.ID})
	if err != nil {
		panic(err)
	}
}
