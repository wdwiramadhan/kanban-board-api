package main

import (
	"os"

	_handler "github.com/wdwiramadhan/kanban-board-api/app/delivery"
	_repository "github.com/wdwiramadhan/kanban-board-api/app/repository"
	_usecase "github.com/wdwiramadhan/kanban-board-api/app/usecase"

	"github.com/gin-gonic/gin"
	"github.com/wdwiramadhan/kanban-board-api/config"
)

func main(){
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	userRepository := _repository.NewUserRepository(db)
	userUsecase := _usecase.NewUserUsecase(userRepository)

	api := router.Group("/")
	_handler.NewUserHanlder(api, userUsecase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":"+ port)
}