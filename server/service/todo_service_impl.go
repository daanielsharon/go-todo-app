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
	DB             *sql.DB
	timeout        time.Duration
}

func NewTodoService(repository repository.TodoRepository, db *sql.DB) TodoService {
	return &TodoServiceImpl{
		TodoRepository: repository,
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
