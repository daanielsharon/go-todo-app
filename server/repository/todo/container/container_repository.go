package containerrepo

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type ContainerRepository interface {
	InitGroup(ctx context.Context, db *sql.DB, userId int) error
	FindGroup(ctx context.Context, db *sql.DB, todo *domain.TodoGroup) (*domain.TodoGroup, error)
	FindTotal(ctx context.Context, db *sql.DB, container *domain.Container) *uint8
	FindById(ctx context.Context, db *sql.DB, id *int64) (*domain.TodoGroup, error)
	Save(ctx context.Context, db *sql.DB, container *domain.Container) *domain.Container
	UpdatePriority(ctx context.Context, db *sql.DB, container *domain.TodoPriority) *domain.TodoPriority
	Delete(ctx context.Context, db *sql.DB, container *domain.Container)
}
