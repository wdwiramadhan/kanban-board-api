package repository

import (
	"context"
	"time"

	"github.com/wdwiramadhan/kanban-board-api/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(Conn *gorm.DB) domain.TaskRepository {
	return &TaskRepository{Conn}
}

func (t *TaskRepository) GetTasks(ctx context.Context) (interface{}, error) {

	type User struct {
		ID       int64  `json:"id" gorm:"primaryKey;"`
		FullName string `json:"full_name" `
		Email    string `json:"email" `
	}
	type Task struct {
		ID          int64     `json:"id" gorm:"primaryKey;"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      bool      `json:"status"`
		UserID      int64     `json:"user_id"`
		CategoryID  int64     `json:"category_id"`
		CreatedAt   time.Time `json:"created_at"`

		User User `json:"user"`
	}

	var tasks []Task
	err := t.Conn.Unscoped().Joins("User").Find(&tasks).Error
	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}

func (t *TaskRepository) StoreTask(ctx context.Context, task *domain.Task) (taskId int64, err error) {
	err = t.Conn.Create(task).Error
	if err != nil {
		return
	}
	taskId = task.ID
	return
}

func (t *TaskRepository) GetTaskByID(ctx context.Context, id int64) (domain.Task, error) {
	var task domain.Task
	err := t.Conn.First(&task, "id = ?", id).Error
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (t *TaskRepository) UpdateTask(ctx context.Context, task *domain.Task) error {
	err := t.Conn.Model(task).Updates(task).Error
	return err
}

func (t *TaskRepository) DeleteTask(ctx context.Context, id int64) error {
	var task domain.Task
	task.ID = id
	err := t.Conn.Delete(&task).Error
	return err
}
