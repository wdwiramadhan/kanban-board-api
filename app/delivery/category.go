package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/wdwiramadhan/kanban-board-api/app/delivery/middleware"
	"github.com/wdwiramadhan/kanban-board-api/app/helper"
	"github.com/wdwiramadhan/kanban-board-api/domain"
)

type CategoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(r *gin.RouterGroup, categoryUsecase domain.CategoryUsecase){
	handler := &CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}
	categoryRoute := r.Group("/categories")
	categoryRoute.Use(middleware.Authentication())
	categoryRoute.GET("/", handler.GetCategories)
	categoryRoute.Use(middleware.Authorization([]string{"admin"}))
	categoryRoute.POST("/", handler.StoreCategory)
	categoryRoute.PATCH(":categoryId", handler.UpdateCategory)
	categoryRoute.DELETE(":categoryId", handler.DeleteCategory)
}

func(c *CategoryHandler) GetCategories(ctx *gin.Context){
	categories, err := c.CategoryUsecase.GetCategories(ctx.Request.Context())
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data" : categories} )
}

func(c *CategoryHandler) StoreCategory(ctx *gin.Context){
	type StoreCategory struct {
		Type string `json:"type" validate:"required"`
	}
	var storeCategory StoreCategory
	err := ctx.ShouldBindJSON(&storeCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = helper.ValidateStruct(storeCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var category domain.Category
	copier.Copy(&category, &storeCategory)
	categoryData, err := c.CategoryUsecase.StoreCategory(ctx.Request.Context(), &category)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"code" : http.StatusCreated, 
		"data" : gin.H{
			"id" : categoryData.ID,
			"type" : categoryData.Type,
			"created_at" : categoryData.CreatedAt,
		},
	})
}

func(c *CategoryHandler) UpdateCategory(ctx *gin.Context){
	type UpdateCategory struct {
		Type string `json:"type" validate:"required"`
	}
	var updateCategory UpdateCategory
	err := ctx.ShouldBindJSON(&updateCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = helper.ValidateStruct(updateCategory)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var category domain.Category
	copier.Copy(&category, &updateCategory)
	categoryId,_ := strconv.ParseInt(ctx.Param("categoryId"), 10, 64)
	category.ID = categoryId
	categoryData, err := c.CategoryUsecase.UpdateCategory(ctx, &category)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK, 
		"data" : gin.H{
			"id" : categoryData.ID,
			"type" : categoryData.Type,
			"updated_at" : categoryData.UpdatedAt,
		},
	})

}

func(c *CategoryHandler) DeleteCategory(ctx *gin.Context){
	categoryId,_ := strconv.ParseInt(ctx.Param("categoryId"), 10, 64)
	err := c.CategoryUsecase.DeleteCategory(ctx, categoryId)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "category has been successfully deleted"})
}