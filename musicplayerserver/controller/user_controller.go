package controller

import (
	"musicplayerserver/model"
	"musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userservice *service.UserService
}

func NewUserController() *UserController {
	uss := service.NewUserService()
	return &UserController{
		userservice: uss,
	}
}

func (usc *UserController) AddUserHandler(c *gin.Context) model.UserInfo {
	username := c.PostForm("username")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := model.UserInfo{
		Username: username,
		Gender:   gender,
		Age:      age,
		Email:    email,
		Password: password,
	}
	usc.userservice.UserRegister(&user)
	return user
}
