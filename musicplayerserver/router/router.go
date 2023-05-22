package router

import (
	"music-player/musicplayerserver/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Posts(r *gin.Engine) {
	r.POST("/register", func(c *gin.Context) {
		user := controller.NewUserController().AddUserHandler(c)
		c.JSON(http.StatusOK, user)
	})
	r.POST("/song/url", func(c *gin.Context) {
		url := controller.NewSongController().GetSongURLHandler(c)
		c.JSON(http.StatusOK, url)
	})
}

func Gets(r *gin.Engine) {
	r.GET("/gettest", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is a test.")
	})
}
