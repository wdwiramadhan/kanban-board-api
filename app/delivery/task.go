package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wdwiramadhan/kanban-board-api/app/delivery/middleware"
	"github.com/wdwiramadhan/kanban-board-api/domain"
)

type TaskHandler struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskHandler(r *gin.RouterGroup, taskUsecase domain.TaskUsecase){
	handler := &TaskHandler{taskUsecase}
	taskRoute := r.Group("/tasks")
	taskRoute.Use(middleware.Authentication())
	taskRoute.GET("/", handler.GetTasks)
	taskRoute.POST("/", handler.StoreTask)
	taskRoute.PATCH("/update-status/:taskId", handler.UpdateStatusTask)
	taskRoute.PATCH("/update-category/:taskId", handler.UpdateCategoryTask)
	taskRoute.DELETE(":taskId", handler.DeleteTask)
}

func(t *TaskHandler) GetTasks(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"code":http.StatusOK})
}

func(t *TaskHandler) StoreTask(ctx *gin.Context){
	ctx.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated})
}

func(t *TaskHandler) UpdateStatusTask(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func(t *TaskHandler) UpdateCategoryTask(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

func(t *TaskHandler) DeleteTask(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}