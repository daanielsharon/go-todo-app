package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type UserRepository interface {
	SaveUsername(ctx context.Context, db *sql.DB, user *domain.User) (*domain.User, error)
	FindUsername(ctx context.Context, db *sql.DB, username *domain.User) (*domain.User, error)
}
