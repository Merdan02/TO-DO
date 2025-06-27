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
	GetByID(ctx context.Context, id int) (*models.Tasks, error)
	UpdateTask(ctx context.Context, task *models.Tasks) error
	DeleteTask(ctx context.Context, id int) error
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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

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

func (repo *TaskRepo) UpdateTask(ctx context.Context, task *models.Tasks) error {
	query := "UPDATE tasks SET title=$1, description=$2, done=$3 WHERE id=$4"
	_, err := repo.db.ExecContext(ctx, query, task.Title, task.Description, task.Done, task.ID)
	if err != nil {
		log.Println("error updating task: ", err)
		return err
	}
	return nil
}

func (repo *TaskRepo) GetByID(ctx context.Context, id int) (*models.Tasks, error) {
	query := "SELECT id, user_id, title, description, done, created_at, updated_at FROM tasks WHERE id=$1"
	row := repo.db.QueryRowContext(ctx, query, id)

	var task models.Tasks
	err := row.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Println("error getting task: ", err)
		return nil, err
	}
	return &task, nil
}

func (repo *TaskRepo) DeleteTask(ctx context.Context, id int) error {
	query := "DELETE FROM tasks WHERE id=$1"
	_, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("error deleting task: ", err)
		return err
	}
	return nil
}
