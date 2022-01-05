package domain

import (
	"context"
	"time"
)

type Task struct {
	ID int64 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title string `json:"title" gorm:"notNull"`
	Description string `json:"description" gorm:"type:text;notNull"`
	Status bool `json:"status" gorm:"notNull"`
	UserID int64 `json:"user_id" gorm:"notNull"`
	CategoryID int64 `json:"category_id" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
	UpdatedAt time.Time `json:"updated_at" gorm:"notNull"`
}

type TaskUsecase interface {
	GetTasks(ctx context.Context) (interface{}, error)
	StoreTask(ctx context.Context, task *Task) (Task, error)
	GetTaskByID(ctx context.Context, id int64) (Task, error)
	UpdateStatusTask(ctx context.Context, task *Task) (Task, error)
	UpdateCategoryTask(ctx context.Context, task *Task) (Task, error)
	DeleteTask(ctx context.Context, id int64) (error)
}

type TaskRepository interface{
	GetTasks(ctx context.Context) (interface{}, error)
	StoreTask(ctx context.Context, task *Task) (taskId int64, err error)
	GetTaskByID(ctx context.Context, id int64) (Task, error)
	UpdateTask(ctx context.Context, task *Task) (error)
	DeleteTask(ctx context.Context, id int64) (error)
}