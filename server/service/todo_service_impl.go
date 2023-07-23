package service

import (
	"context"
	"database/sql"
	"fmt"
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

func (s *TodoServiceImpl) CreateTodo(c context.Context, req *web.TodoCreateRequest) *web.TodoCreateResponse {
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

	r, err := s.TodoRepository.SaveTodo(ctx, s.DB, newUser)
	if err != nil {
		panic(err)
	}

	res := &web.TodoCreateResponse{
		ID:      r.ID,
		Name:    r.Name,
		UserID:  r.UserID,
		GroupID: r.GroupID,
	}

	return res
}

func (s *TodoServiceImpl) RemoveTodo(c context.Context, req *web.TodoDeleteRequest) {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	_, err := s.TodoRepository.FindTodoById(ctx, s.DB, int(req.ID))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = s.TodoRepository.DeleteTodo(ctx, s.DB, &domain.TodoList{ID: req.ID})
	if err != nil {
		panic(err)
	}
}

func (s *TodoServiceImpl) GetTodoByUsername(c context.Context, req *web.TodoGetRequest) *[]web.TodoGetResponse {
	ctx, cancel := context.WithTimeout(c, s.Timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	_, err := s.UserRepository.FindUsername(ctx, s.DB, &user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	res, err := s.TodoRepository.FindTodoByUsername(ctx, s.DB, &user)
	helper.PanicIfError(err)

	var response []web.TodoGetResponse

	for _, val := range *res {
		var item interface{}

		fmt.Println("val", val)

		if val.Item[0] == nil {
			item = make([]interface{}, 0)
		} else {
			item = val.Item
		}

		// for _, each := range val {
		// 	fmt.Println("each", each)
		// 	if each != nil {
		// 		item = append(item, interface{})
		// 		continue
		// 	}
		// 	item = append(item, each)
		// }

		response = append(response, web.TodoGetResponse{
			ID:        val.ID,
			GroupName: val.Name,
			Item:      item,
			Priority:  val.Priority,
		})
	}

	return &response
}
