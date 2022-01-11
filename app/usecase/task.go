package usecase

import (
	"context"
	"time"

	"github.com/wdwiramadhan/kanban-board-api/domain"
)

type TaskUsecase struct {
	taskRepository     domain.TaskRepository
	categoryRepository domain.CategoryRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository, categoryRepository domain.CategoryRepository) domain.TaskUsecase {
	return &TaskUsecase{taskRepository, categoryRepository}
}

func (t *TaskUsecase) GetTasks(ctx context.Context) (interface{}, error) {
	task, err := t.taskRepository.GetTasks(ctx)
	if err != nil {
		return []domain.Task{}, domain.ErrInternalServerError
	}
	return task, nil
}

func (t *TaskUsecase) StoreTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := t.categoryRepository.GetCategoryByID(ctx, task.CategoryID)
	if err != nil {
		return domain.Task{}, domain.ErrInternalServerError
	}
	taskId, err := t.taskRepository.StoreTask(ctx, task)
	if err != nil {
		return domain.Task{}, domain.ErrInternalServerError
	}
	task.ID = taskId
	return *task, nil
}

func (t *TaskUsecase) GetTaskByID(ctx context.Context, id int64) (domain.Task, error) {
	task, err := t.taskRepository.GetTaskByID(ctx, id)
	if err != nil {
		return domain.Task{}, domain.ErrNotFound
	}
	return task, nil

}

func (t *TaskUsecase) UpdateTask(ctx context.Context, task *domain.Task) (domain.Task, error) {
	_, err := t.taskRepository.GetTaskByID(ctx, task.ID)
	if err != nil {
		return domain.Task{}, domain.ErrNotFound
	}
	task.UpdatedAt = time.Now()
	err = t.taskRepository.UpdateTask(ctx, task)
	if err != nil {
		return domain.Task{}, domain.ErrInternalServerError
	}
	return *task, nil
}

func (t *TaskUsecase) DeleteTask(ctx context.Context, id int64) error {
	_, err := t.taskRepository.GetTaskByID(ctx, id)
	if err != nil {
		return domain.ErrNotFound
	}
	err = t.taskRepository.DeleteTask(ctx, id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}
