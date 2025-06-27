package repository

import (
	"context"
	"database/sql"
	"log"
	"todo-app/internal/models"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) TaskRepository {
	return &TaskRepo{
		db: db,
	}
}

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Tasks) error
	GetAllTasks(ctx context.Context) ([]models.Tasks, error)
}

func (repo *TaskRepo) CreateTask(ctx context.Context, task *models.Tasks) error {
	query := "INSERT INTO tasks (user_id, title, description, done ) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	err := repo.db.QueryRowContext(ctx, query, task.UserID, task.Title, task.Description, task.Done).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (repo *TaskRepo) GetAllTasks(ctx context.Context) ([]models.Tasks, error) {
	query := "SELECT id, user_id, title, description, done, created_at, updated_at FROM tasks"
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("error getting all tasks: ", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Tasks
	for rows.Next() {
		var task models.Tasks
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("error getting all tasks: ", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Println("error getting all tasks: ", err)
		return nil, err
	}
	return tasks, nil
}
