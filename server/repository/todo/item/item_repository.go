package itemrepo

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type ItemRepository interface {
	Save(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) (*domain.TodoListInsertUpdate, error)
	FindByUsername(ctx context.Context, db *sql.DB, user *domain.User) (*[]domain.Todo, error)
	FindById(ctx context.Context, db *sql.DB, id int64) (*domain.TodoList, error)
	FindByName(ctx context.Context, db *sql.DB, name string) (*domain.TodoList, error)
	Update(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) *domain.TodoListInsertUpdate
	Delete(ctx context.Context, db *sql.DB, todo *domain.TodoList) error
}
