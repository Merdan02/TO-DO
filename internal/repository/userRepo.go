package repository

import (
	"context"
	"database/sql"
	"log"
	"todo-app/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *models.UserModel) error {
	query := `INSERT INTO users(username, password_hash, role) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := repo.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Role).
		Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}
	return nil
}
