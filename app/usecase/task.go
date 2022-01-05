package usecase

import (
	"context"

	"github.com/wdwiramadhan/kanban-board-api/domain"
)

type TaskUsecase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{taskRepository}
}

func(t *TaskUsecase) GetTasks(ctx context.Context) (interface{}, error) {
	return domain.Task{}, nil
}

func(t *TaskUsecase) StoreTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func(t *TaskUsecase) GetTaskByID(ctx context.Context, id int64) (domain.Task, error) {
	return domain.Task{}, nil
}

func(t *TaskUsecase) UpdateStatusTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func(t *TaskUsecase) UpdateCategoryTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
	return domain.Task{}, nil
}

func(t *TaskUsecase) DeleteTask(ctx context.Context, id int64) (error) {
	return  nil
}
