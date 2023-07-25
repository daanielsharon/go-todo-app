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

func (u *UserRepositoryImpl) FindUsername(ctx context.Context, db *sql.DB, username *domain.User) (*domain.User, error) {
	query := "SELECT id, username FROM users WHERE username = $1"
	row, err := db.QueryContext(ctx, query, username.Username)

	if err != nil {
		return &domain.User{}, err
	}

	user := domain.User{}
	if row.Next() {
		err := row.Scan(&user.ID, &user.Username)
		if err != nil {
			return &user, err
		}
		return &user, nil
	} else {
		return &user, errors.New("user is not found")
	}
}

func (u *UserRepositoryImpl) SaveUsername(ctx context.Context, db *sql.DB, user *domain.User) (*domain.User, error) {
	var userId int64
	query := "INSERT INTO users(username) VALUES($1) RETURNING id"
	err := db.QueryRowContext(ctx, query, user.Username).Scan(&userId)

	if err != nil {
		return &domain.User{}, err
	}

	user.ID = userId
	return user, nil
}
