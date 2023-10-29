package controllers

import (
	"REST-API/models"
	"REST-API/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {

	return UserController{UserService: userService}
}

func (uc *UserController) CreatUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.CreatUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Created"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	name := ctx.Param("name")
	user, err := uc.UserService.GetUser(&name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Created"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	name := ctx.Param("name")
	err := uc.UserService.DeleteUser(&name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Deleted"})

}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/users")
	router.POST("/", uc.CreatUser)
	router.GET("/:name", uc.GetUser)
	router.GET("/all", uc.GetAllUsers)
	router.PUT("/", uc.UpdateUser)
	router.DELETE("/:name", uc.DeleteUser)
}
