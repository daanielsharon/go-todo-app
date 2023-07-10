package service

import (
	"context"
	"database/sql"
	"server/model/domain"
	"server/model/web"
	"server/repository"
	"time"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	UserRepository repository.UserRepository
	DB             *sql.DB
	timeout        time.Duration
}

func NewTodoService(todoRepository repository.TodoRepository, userRepository repository.UserRepository, db *sql.DB) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		UserRepository: userRepository,
		DB:             db,
		timeout:        time.Duration(2) * time.Second,
	}
}

func (s *TodoServiceImpl) CreateTodo(c context.Context, req *web.TodoCreateRequest) (*web.TodoCreateResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	newUser := &domain.TodoListInsertUpdate{
		Name:    req.Name,
		GroupID: req.GroupID,
		UserID:  req.UserID,
	}

	r, err := s.TodoRepository.SaveTodo(ctx, s.DB, newUser)
	if err != nil {
		return nil, err
	}

	res := &web.TodoCreateResponse{
		ID:      r.ID,
		Name:    r.Name,
		UserID:  r.UserID,
		GroupID: r.GroupID,
	}

	return res, nil
}

func (s *TodoServiceImpl) RemoveTodo(c context.Context, req *web.TodoDeleteRequest) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	_, err := s.TodoRepository.FindTodoById(ctx, s.DB, int(req.ID))
	if err != nil {
		return err
	}

	err = s.TodoRepository.DeleteTodo(ctx, s.DB, &domain.TodoList{ID: req.ID})
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoServiceImpl) GetTodoByUsername(c context.Context, req *web.TodoGetRequest) (*[]web.TodoGetResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user := domain.User{
		Username: req.Username,
	}

	_, err := s.UserRepository.FindUsername(ctx, s.DB, &user)
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
