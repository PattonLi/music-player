package main

import (
	"musicplayerserver/dao"
	"musicplayerserver/router"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.Init()
	/*
		userdao := dao.UserDao{}
		user := model.UserInfo{
			Username: "JohnDoe",
			Gender:   "Male",
			Age:      "30",
			Email:    "johndoe@example.com",
			Password: "password",
		}
		userdao.AddUser(&user)
	*/
	r := gin.Default()
	router.Posts(r)
	r.Run()
}
