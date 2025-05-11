package repository

import (
	"context"
	"database/sql"
	"github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (ur *UsersRepository) CreateUser(ctx context.Context, newUser *domain.User) error {
	const q = `
		INSERT INTO users(login, password_hash)
		VALUES (?, ?)
	`

	_, err := ur.db.ExecContext(ctx, q, newUser.Login, newUser.PasswordHash)
	return err
}

func (ur *UsersRepository) GetUserByLogin(ctx context.Context, login string) (*domain.User, error) {
	const q = `
		SELECT id, login, password_hash, created_at 
		FROM users 
		WHERE login = ?
	`

	row := ur.db.QueryRowContext(ctx, q, login)
	var user domain.User
	err := row.Scan(&user.Id, &user.Login, &user.PasswordHash, &user.CreationTime)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UsersRepository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	const q = `
		SELECT id, login, password_hash, created_at
		FROM users
		WHERE id = ?
	`

	row := ur.db.QueryRowContext(ctx, q, id)
	var user domain.User
	err := row.Scan(&user.Id, &user.Login, &user.PasswordHash, &user.CreationTime)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
