package repository

import (
	"context"

	"github.com/wdwiramadhan/kanban-board-api/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(Conn *gorm.DB) domain.TaskRepository {
	return &TaskRepository{Conn}
}

func(t *TaskRepository) GetTasks(ctx context.Context) (interface{}, error) {
	return domain.Task{}, nil
}

func(t *TaskRepository) StoreTask(ctx context.Context, task *domain.Task) (taskId int64, err error) {
	return 0, nil
}

func(t *TaskRepository) GetTaskByID(ctx context.Context, id int64) (domain.Task, error) {
	return domain.Task{}, nil
}

func(t *TaskRepository) UpdateTask(ctx context.Context, task *domain.Task) (error) {
	return nil
}

func(t *TaskRepository) DeleteTask(ctx context.Context, id int64) (error){
	return nil
}
