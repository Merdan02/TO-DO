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

func NewUserRepo(db *sql.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
	GetAllUsers(ctx context.Context) ([]models.UserModel, error)
	GetUserByID(ctx context.Context, userID int) (*models.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error)
	UpdateUser(ctx context.Context, user *models.UserModel) error
	DeleteUser(ctx context.Context, userID int) error
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
func (repo *UserRepo) GetAllUsers(ctx context.Context) ([]models.UserModel, error) {
	query := `SELECT id, username, password_hash, role, created_at, updated_at FROM users`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
			return
		}
	}(rows)
	var users []models.UserModel
	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("Error getting all users: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) GetUserByID(ctx context.Context, id int) (*models.UserModel, error) {
	query := "SELECT id, username, password_hash, role, created_at, updated_at FROM users WHERE id = $1"
	row := repo.db.QueryRowContext(ctx, query, id)

	var user models.UserModel
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepo) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	user := &models.UserModel{}
	query := "SELECT id, username, password_hash, role, created_at, updated_at FROM users WHERE username = $1"
	err := repo.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) UpdateUser(ctx context.Context, user *models.UserModel) error {
	query := "UPDATE users SET username = $1, password_hash = $2, role = $3 WHERE id = $4"
	_, err := repo.db.ExecContext(ctx, query, user.Username, user.Password, user.Role, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func (repo *UserRepo) DeleteUser(ctx context.Context, userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := repo.db.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
