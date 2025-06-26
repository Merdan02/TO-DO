package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"todo-app/internal/models"
	"todo-app/internal/repository"
)

type TaskServ struct {
	repo repository.TaskRepository
	Log  *zap.Logger
}

func NewTaskServ(repo repository.TaskRepository, logger *zap.Logger) TaskService {
	return &TaskServ{
		repo: repo,
		Log:  logger,
	}
}

type TaskService interface {
	CreateTask(ctx context.Context, task *models.Tasks) error
}

func (s *TaskServ) CreateTask(ctx context.Context, task *models.Tasks) error {
	if task.Title == "" || task.Description == "" {
		s.Log.Error("task title or description is empty")
		return fmt.Errorf("task title or description is empty")
	}

	if err := s.repo.CreateTask(ctx, task); err != nil {
		s.Log.Error("create task failed", zap.Error(err))
		return err
	}
	return nil
}
