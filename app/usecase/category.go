package usecase

import (
	"context"
	"time"

	"github.com/wdwiramadhan/kanban-board-api/domain"
)

type CategoryUsecase struct {
	categoryRepository domain.CategoryRepository
}

func NewCategoryUsecase(c domain.CategoryRepository) domain.CategoryUsecase {
	return &CategoryUsecase{categoryRepository: c}
}

func(c *CategoryUsecase) GetCategories(ctx context.Context) (interface{}, error) {
	categories, err := c.categoryRepository.GetCategories(ctx)
	if err != nil {
		return []domain.Category{}, domain.ErrInternalServerError
	}
	return categories, nil
}

func(c *CategoryUsecase) StoreCategory(ctx context.Context, category *domain.Category) (domain.Category, error) {
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	categoryId, err := c.categoryRepository.StoreCategory(ctx, category)
	if err != nil {
		return domain.Category{}, domain.ErrInternalServerError
	}
	category.ID = categoryId
	return *category, nil
}

func(c *CategoryUsecase) GetCategoryByID(ctx context.Context, id int64) (domain.Category, error) {
	category, err := c.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return domain.Category{}, domain.ErrNotFound
	}
	return category, nil
}

func(c *CategoryUsecase) UpdateCategory(ctx context.Context, category *domain.Category) (domain.Category, error) {
	_, err := c.categoryRepository.GetCategoryByID(ctx, category.ID)
	if err != nil {
		return domain.Category{}, domain.ErrNotFound
	}
	category.UpdatedAt = time.Now()
	err = c.categoryRepository.UpdateCategory(ctx, category)
	if err != nil {
		return domain.Category{}, domain.ErrInternalServerError
	}
	return *category, nil
}

func(c *CategoryUsecase) DeleteCategory(ctx context.Context, id int64) (error) {
	_, err := c.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return  domain.ErrNotFound
	}
	err = c.categoryRepository.DeleteCategory(ctx, id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}
	
	
	
	

