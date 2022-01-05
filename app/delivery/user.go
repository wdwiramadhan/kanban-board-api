package delivery

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/wdwiramadhan/kanban-board-api/app/delivery/middleware"
	"github.com/wdwiramadhan/kanban-board-api/app/helper"
	"github.com/wdwiramadhan/kanban-board-api/domain"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHanlder(r *gin.RouterGroup, userUsecase domain.UserUsecase){
	handler := &UserHandler{userUsecase}
	r.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	userRoute := r.Group("/users")
	userRoute.POST("/register", handler.Register)
	userRoute.POST("/login", handler.Login)
	userRoute.Use(middleware.Authentication())
	userRoute.PUT("/update-account", handler.UpdateAccount)
	userRoute.DELETE("/delete-account", handler.DeleteAccount)
} 

func (u *UserHandler) Register(ctx *gin.Context){
	type UserRegister struct{
		FullName string `json:"full_name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=16"`
	}
	var userRegister UserRegister
	err := ctx.ShouldBindJSON(&userRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = helper.ValidateStruct(userRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var user domain.User
	copier.Copy(&user, &userRegister)
	userData, err := u.userUsecase.Register(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"code" : http.StatusCreated, 
		"data" : gin.H{
			"id" : userData.ID,
			"full_name" : userData.FullName,
			"email" : userData.Email,
			"created_at" : userData.CreatedAt,
		},
	})
}

func(u *UserHandler) Login(ctx *gin.Context){
	type UserLogin struct{
		Email string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	var userLogin UserLogin
	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = helper.ValidateStruct(userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var user domain.User
	copier.Copy(&user, &userLogin)
	token, err := u.userUsecase.Login(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK, 
		"data" : gin.H {
			"token" : token,
		},
	})

}

func(u *UserHandler) UpdateAccount(ctx *gin.Context){
	type UserUpdate struct {
		FullName string `json:"full_name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
	var userUpdate UserUpdate
	err := ctx.ShouldBindJSON(&userUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = helper.ValidateStruct(userUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var user domain.User
	copier.Copy(&user, &userUpdate)
	userAuth := ctx.MustGet("user").(jwt.MapClaims)
	userID :=  int64(userAuth["id"].(float64))
	user.ID = userID
	userData, err := u.userUsecase.UpdateUser(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H {
			"id" : userData.ID,
			"full_name" : userData.FullName,
			"email" : userData.Email,
			"updated_at" : userData.UpdatedAt,
		},
	})
}

func(u *UserHandler) DeleteAccount(ctx *gin.Context){
	userAuth := ctx.MustGet("user").(jwt.MapClaims)
	userID :=  int64(userAuth["id"].(float64))
	err := u.userUsecase.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(getStatusCode(err), gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "your account has been successfully deleted"})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
		case domain.ErrInternalServerError:
			return http.StatusInternalServerError
		case domain.ErrNotFound:
			return http.StatusNotFound
		case domain.ErrConflict:
			return http.StatusConflict
		case domain.ErrUnauthorized:
			return http.StatusUnauthorized
		default:
			return http.StatusInternalServerError
	}
}