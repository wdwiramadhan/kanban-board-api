package repository

import (
	"context"
	"time"

	"github.com/wdwiramadhan/kanban-board-api/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(Conn *gorm.DB) domain.CategoryRepository {
	return &CategoryRepository{Conn}
}

func (c *CategoryRepository) GetCategories(ctx context.Context) (interface{}, error) {

	type Task struct {
		ID          int64     `json:"id" gorm:"primaryKey;autoIncrement:true"`
		Title       string    `json:"title" gorm:"notNull"`
		Description string    `json:"description" gorm:"type:text;notNull"`
		UserID      int64     `json:"user_id" gorm:"notNull"`
		CategoryID  int64     `json:"category_id" gorm:"notNull"`
		CreatedAt   time.Time `json:"created_at" gorm:"notNull"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"notNull"`
	}

	type Category struct {
		domain.Category
		Tasks []Task
	}
	var categoriesTasks []Category
	err := c.Conn.Preload("Tasks").Find(&categoriesTasks).Error
	if err != nil {
		return []domain.Category{}, err
	}
	return categoriesTasks, nil
}

func (c *CategoryRepository) StoreCategory(ctx context.Context, category *domain.Category) (categoryId int64, err error) {
	err = c.Conn.Create(category).Error
	if err != nil {
		return
	}
	categoryId = category.ID
	return
}

func (c *CategoryRepository) GetCategoryByID(ctx context.Context, id int64) (domain.Category, error) {
	var category domain.Category
	err := c.Conn.First(&category, "id=?", id).Error
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (c *CategoryRepository) UpdateCategory(ctx context.Context, category *domain.Category) error {
	err := c.Conn.Model(category).Updates(category).Error
	return err
}

func (c *CategoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	var category domain.Category
	category.ID = id
	err := c.Conn.Delete(&category).Error
	return err
}
