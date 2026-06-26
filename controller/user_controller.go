package controller

import (
	"net/http"
	"strconv"

	"users-api/model"
	"users-api/usecase"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUseCase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) userController {
	return userController{
		userUseCase: usecase,
	}
}

func (uc *userController) GetUsers(ctx *gin.Context) {
	users, err := uc.userUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	user := model.NewUser("", 0)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = uc.userUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uc *userController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "The user id can't be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "The user id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userUseCase.GetUserByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		response := model.Response{
			Message: "User not found on database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	user := model.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser, err := uc.userUseCase.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if updatedUser == nil {
		response := model.Response{
			Message: "User not found on database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (uc *userController) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "The user id can't be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "The user id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userUseCase.DeleteUserByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		response := model.Response{
			Message: "User not found on database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
