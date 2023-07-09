package repository

import (
	"context"
	"database/sql"
	"errors"
	"server/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) GetUsername(ctx context.Context, username *domain.User, db *sql.DB) (*domain.User, error) {
	query := "SELECT id, username FROM users WHERE username = %1"
	rows, err := db.QueryContext(ctx, query, username.Username)

	if err != nil {
		return &domain.User{}, err
	}

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return &user, err
		}
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}
