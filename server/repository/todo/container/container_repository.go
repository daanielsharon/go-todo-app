package containerrepo

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type ContainerRepository interface {
	InitGroup(ctx context.Context, db *sql.DB, userId int) error
	FindGroup(ctx context.Context, db *sql.DB, todo *domain.TodoGroup) (*domain.TodoGroup, error)
	UpdatePriority(ctx context.Context, db *sql.DB, container *domain.TodoPriority) *domain.TodoPriority
}
