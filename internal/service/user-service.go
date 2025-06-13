package service

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
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
