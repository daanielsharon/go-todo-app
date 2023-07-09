package repository

import (
	"context"
	"database/sql"
	"server/model/domain"
)

type UserRepository interface {
	GetUsername(ctx context.Context, username *domain.User, db *sql.DB) (*domain.User, error)
}
