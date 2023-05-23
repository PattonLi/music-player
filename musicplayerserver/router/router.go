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
}

func Gets(r *gin.Engine) {
	r.GET("/lyric", func(c *gin.Context) {
		lyric := controller.NewSongController().GetSongHandler(c)
		c.String(http.StatusOK, lyric)
	})

}
