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
	GetAllTasks(ctx context.Context) ([]models.Tasks, error)
	GetByID(ctx context.Context, id int) (*models.Tasks, error)
	UpdateTask(ctx context.Context, task *models.Tasks) error
	DeleteTask(ctx context.Context, id int) error
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

func (s *TaskServ) GetAllTasks(ctx context.Context) ([]models.Tasks, error) {
	tasks, err := s.repo.GetAllTasks(ctx)
	if err != nil {
		s.Log.Error("error getting all tasks", zap.Error(err))
	}
	return tasks, err
}

func (s *TaskServ) UpdateTask(ctx context.Context, task *models.Tasks) error {
	if task.ID == 0 {
		s.Log.Error("task id is empty")
		return fmt.Errorf("task id is empty")
	}
	if err := s.repo.UpdateTask(ctx, task); err != nil {
		s.Log.Error("update task failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *TaskServ) GetByID(ctx context.Context, id int) (*models.Tasks, error) {
	if id == 0 {
		s.Log.Error("task id is empty")
		return nil, fmt.Errorf("task id is empty")
	}

	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.Log.Error("get task failed", zap.Error(err))
		return nil, err
	}
	return task, nil
}

func (s *TaskServ) DeleteTask(ctx context.Context, id int) error {
	if id == 0 {
		s.Log.Error("task id is empty")
		return fmt.Errorf("task id is empty")
	}
	if err := s.repo.DeleteTask(ctx, id); err != nil {
		s.Log.Error("delete task failed", zap.Error(err))
		return err
	}
	return nil
}
