package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type UserRepository interface {
	SaveUsername(ctx context.Context, db *sql.DB, user *domain.User) (*domain.User, error)
	GetUsername(ctx context.Context, username *domain.User, db *sql.DB) (*domain.User, error)
}
