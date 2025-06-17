package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"todo-app/internal/models"
	"todo-app/internal/repository"
)

type UserServ struct {
	repo repository.UserRepository
	Log  *zap.Logger
}

func NewUserServ(repo repository.UserRepository, logger *zap.Logger) UserService {
	return &UserServ{
		repo: repo,
		Log:  logger,
	}
}

type UserService interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
	GetAllUsers(ctx context.Context) ([]models.UserModel, error)
	GetUserByID(ctx context.Context, userID int) (*models.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error)
	UpdateUser(ctx context.Context, user *models.UserModel) error
	DeleteUser(ctx context.Context, userID int) error
}

func (s *UserServ) CreateUser(ctx context.Context, user *models.UserModel) error {
	if user.Username == "" || user.Password == "" {
		s.Log.Error("username or password  is empty", zap.String("username", user.Username))
		return errors.New("username or password  is empty")
	}

	if user.Role != "admin" && user.Role != "user" {
		s.Log.Error("role is invalid", zap.String("role", user.Role))
		return errors.New("role is invalid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Error("hashing password failed", zap.Error(err))
		return errors.New("hashing password failed")
	}
	user.Password = string(hashedPassword)

	if err := s.repo.CreateUser(ctx, user); err != nil {
		s.Log.Error("failed to create user", zap.String("username", user.Username), zap.Error(err))
		return err
	}
	return nil
}

func (s *UserServ) GetAllUsers(ctx context.Context) ([]models.UserModel, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		s.Log.Error("failed to get all users", zap.Error(err))
		return nil, errors.New("failed to get all users")
	}
	return users, err

}

func (s *UserServ) GetUserByID(ctx context.Context, userID int) (*models.UserModel, error) {
	if userID <= 0 {
		s.Log.Error("user id is empty")
		return nil, errors.New("user id is empty")
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		s.Log.Error("failed to get user", zap.Int("user_id", userID))
		return nil, errors.New("failed to get user")
	}
	return user, nil
}

func (s *UserServ) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	if username == "" {
		s.Log.Error("username is empty")
		return nil, errors.New("username is empty")
	}
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		s.Log.Error("failed to get user", zap.String("username", username))
		return nil, errors.New("failed to get user")
	}
	return user, nil
}

func (s *UserServ) UpdateUser(ctx context.Context, user *models.UserModel) error {
	if user.Username == "" || user.Password == "" {
		s.Log.Error("username or password  is empty")
		return errors.New("username or password  is empty")
	}
	if user.Role != "admin" && user.Role != "user" {
		s.Log.Error("role is invalid", zap.String("role", user.Role))
		return errors.New("role is invalid")
	}
	if !strings.Contains(user.Password, "$2a$") {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			s.Log.Error("hashing password failed", zap.Error(err))
			return errors.New("hashing password failed")
		}
		user.Password = string(hashedPassword)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Error("hashing password failed", zap.Error(err))
		return errors.New("hashing password failed")
	}
	user.Password = string(hashedPassword)
	if err := s.repo.UpdateUser(ctx, user); err != nil {
		s.Log.Error("failed to update user", zap.String("username", user.Username), zap.Error(err))
		return err
	}
	return nil
}

func (s *UserServ) DeleteUser(ctx context.Context, userID int) error {
	if userID == 0 {
		s.Log.Error("user id is empty")
		return errors.New("user id is empty")
	}
	err := s.repo.DeleteUser(ctx, userID)
	if err != nil {
		s.Log.Error("failed to delete user", zap.Int("user_id", userID))
		return errors.New("failed to delete user")
	}
	return nil
}
