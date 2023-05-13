package controller

import (
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/service"

	"github.com/gin-gonic/gin"
)

var userservice = service.UserService{}

func UserControllerInit() {
	userservice.NewUserService()
}

func AddUserHandler(c *gin.Context) model.UserInfo {
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
	userservice.UserRegister(&user)
	return user
}
