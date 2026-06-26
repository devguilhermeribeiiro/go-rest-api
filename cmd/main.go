package main

import (
	"users-api/controller"
	"users-api/database"
	"users-api/respository"
	"users-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := database.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	UserRepository := respository.NewUserRepository(dbConnection)
	UserUseCase := usecase.NewUserUsecase(UserRepository)
	UserController := controller.NewUserController(UserUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/users", UserController.GetUsers)
	server.POST("/user", UserController.CreateUser)
	server.GET("/user/:id", UserController.GetUserByID)
	server.PUT("/user/update", UserController.UpdateUser)
	server.DELETE("/user/:id", UserController.DeleteUserByID)

	server.Run()
}
