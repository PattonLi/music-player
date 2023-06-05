package main

import (
	"music-player/musicplayerserver/dao"
	"music-player/musicplayerserver/router"
	"music-player/musicplayerserver/utils/oss"
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

	//跨域配置，解决CORS跨域问题
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
	//开启OSS访问
	utils.CreateClient()
	r.Use(cors.New(corsConfig))
	router.Posts(r)
	router.GETs(r)
	r.Run("127.0.0.1:4000")

}
