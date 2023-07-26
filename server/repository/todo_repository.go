package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type TodoRepository interface {
	InitGroup(ctx context.Context, db *sql.DB, userId int) error
	Save(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) (*domain.TodoListInsertUpdate, error)
	FindByUsername(ctx context.Context, db *sql.DB, user *domain.User) (*[]domain.Todo, error)
	FindGroup(ctx context.Context, db *sql.DB, todo *domain.TodoGroup) (*domain.TodoGroup, error)
	FindById(ctx context.Context, db *sql.DB, id int64) (*domain.TodoList, error)
	Update(ctx context.Context, db *sql.DB, todo *domain.TodoListInsertUpdate) *domain.TodoListInsertUpdate
	Delete(ctx context.Context, db *sql.DB, todo *domain.TodoList) error
}
