package service

import (
	"context"
	"database/sql"
	"server/exception"
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
	timeout        time.Duration
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, userRepository repository.UserRepository, db *sql.DB, validator *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		UserRepository: userRepository,
		DB:             db,
		timeout:        time.Duration(2) * time.Second,
		Validate:       validator,
	}
}

func (s *TodoServiceImpl) CreateTodo(c context.Context, req *web.TodoCreateRequest) *web.TodoCreateResponse {

	ctx, cancel := context.WithTimeout(c, s.timeout)
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
	ctx, cancel := context.WithTimeout(c, s.timeout)
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

func (s *TodoServiceImpl) GetTodoByUsername(c context.Context, req *web.TodoGetRequest) (*[]web.TodoGetResponse, error) {
	err := s.Validate.Struct(req)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	_, err = s.UserRepository.FindUsername(ctx, s.DB, &user)
	if err != nil {
		return &[]web.TodoGetResponse{}, err
	}

	res, err := s.TodoRepository.FindTodoByUsername(ctx, s.DB, &user)
	if err != nil {
		return &[]web.TodoGetResponse{}, err
	}

	var response []web.TodoGetResponse

	for _, val := range *res {
		var item []web.TodoItemResponse

		for _, each := range val.Item {
			item = append(item, web.TodoItemResponse{
				ID:   each.ID,
				Name: each.Name,
			})
		}

		response = append(response, web.TodoGetResponse{
			ID:       val.ID,
			Name:     val.Name,
			Item:     item,
			Priority: val.Priority,
		})
	}

	return &response, nil
}
