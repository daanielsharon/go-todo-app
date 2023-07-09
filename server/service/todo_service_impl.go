package service

import (
	"context"
	"database/sql"
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

func (s *TodoServiceImpl) SaveTodo(ctx context.Context, req *web.TodoCreateRequest) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
}
