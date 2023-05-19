package main

import (
	"musicplayerserver/dao"
	"musicplayerserver/router"
	"time"

	"github.com/gin-gonic/gin"

	//CORS
	"github.com/gin-contrib/cors"
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
	//跨域配置
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "PATCH",
		},
		AllowHeaders: []string{
			"Content-Type", "Access-Token", "Authorization",
		},
		MaxAge: 6 * time.Hour,
	}
	r.Use(cors.New(corsConfig))
	//
	router.Posts(r)
	r.Run()
}
