package router

import (
	"musicplayerserver/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Posts(r *gin.Engine) {
	r.POST("/register", func(c *gin.Context) {
		user := controller.AddUserHandler(c)
		c.JSON(http.StatusOK, user)
	})
}
