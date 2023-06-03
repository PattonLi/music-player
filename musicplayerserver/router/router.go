package router

import (
	"fmt"
	"music-player/musicplayerserver/controller"
	"net/http"
	"music-player/musicplayerserver/utils"
	"github.com/gin-gonic/gin"
)



func Posts(r *gin.Engine) {

	//鉴权路由，需要鉴权的post api将r改为authorized
	authorized := r.Group("/")
	authorized.Use(utils.AuthMiddleware)
	//添加用户
	r.POST("/User/addInfo", func(c *gin.Context) {
		user, err := controller.NewUserController().AddUserHandler(c)
		var code string
		if err != nil {
			fmt.Print(err.Error())
			code = "300"
		} else {
			code = "200"
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"data": []gin.H{
				{
					"user_id":    user.ID,
					"created_at": user.CreatedAt,
					"updated_at": user.UpdatedAt,
					"deleted_at": user.DeletedAt,
					"username":   user.Username,
					"gender":     user.Gender,
					"age":        user.Age,
					"email":      user.Email,
					"password":   user.Password,
					"nickname":   user.Nickname,
					"phone":      user.Phone,
				}},
		})
	})

	//修改用户信息
	r.POST("/User/modifyInfo", func(c *gin.Context) {

	})

	//手机号登录（未实现）
	r.POST("/login", func(c *gin.Context) {
		username, userID, err := controller.NewUserController().UserLoginHandler(c)
		tokenString := ""
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "token": tokenString, "userid": userID})
		} else {
			tokenString, err = utils.CreateToken(username, "false")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "token": tokenString, "adminid": userID})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "登录成功！", "token": tokenString, "userid": userID})
			}
		}
	})

	//管理员登录（未修改）
	r.POST("/admin/login", func(c *gin.Context) {
		adminname, adminID, err := controller.NewAdminUserController().AdminLoginHandler(c)
		tokenString := ""
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "token": tokenString, "adminid": adminID})
		} else {
			tokenString, err = utils.CreateToken(adminname, "true")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "token": tokenString, "adminid": adminID})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "登录成功！", "token": tokenString, "adminid": adminID})
			}
		}
	})

	//添加歌曲
	r.POST("/admin/addsong", func(c *gin.Context) {
		result := controller.NewSongController().AddSongHandler(c)
		if result {
			c.JSON(http.StatusOK, gin.H{"message": "成功添加歌曲"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "添加歌曲失败"})
		}
	})

}

func GETs(r *gin.Engine) {
	//鉴权路由，需要鉴权的post api将r改为authorized
	authorized := r.Group("/")
	authorized.Use(utils.AuthMiddleware)

	//获得特定页所有用户信息
	r.GET("/User/pageAllInfo", func(c *gin.Context) {
		users,totalPage := controller.NewUserController().AllUserInfoHandler(c)
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"data": users,
			"totalPage":totalPage,
			},
		)
	})

	//获得特定名称用户信息
	r.GET("/User/theInfo", func(c *gin.Context) {
		users,err := controller.NewUserController().UserInfoHandler(c)
		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code": "300",
				"data": users,
				},
			)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": "200",
				"data": users,
				},
			)
		}
	})

	//删除用户信息
	r.GET("/User/deleteInfo", func(c *gin.Context) {

	})

	//获得所有管理员信息
	r.GET("/adminUser/allInfo", func(c *gin.Context) {
		adminusers := controller.NewAdminUserController().AllAdminInfoHandler()
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"data": []gin.H{
				{
					"adminusers": adminusers,
				},
			},
		})
	})

	r.GET("/song/lyric", func(c *gin.Context) {
		lyric, err := controller.NewSongController().GetSongLyricHandler(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "没有此歌曲"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "获取歌词成功", "lyric": lyric})
		}
	})

	r.GET("/song/detail", func(c *gin.Context) {
		songname, singer, err := controller.NewSongController().GetSongDetailHandler(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "没有此歌曲"})
		} else {

			c.JSON(http.StatusOK, gin.H{"message": "获取信息成功", "songname": songname, "singer": singer})
		}
	})
}
