package router

import (
	"fmt"
	"music-player/musicplayerserver/controller"
	"music-player/musicplayerserver/model"
	"music-player/musicplayerserver/utils/jwt"
	"net/http"
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
					"user_id":  user.ID,
					"username": user.Username,
					"gender":   user.Gender,
					"age":      user.Age,
					"email":    user.Email,
					"password": user.Password,
					"nickname": user.Nickname,
					"phone":    user.Phone,
				}},
		})
	})

	//修改用户信息
	r.POST("/User/modifyInfo", func(c *gin.Context) {
		err := controller.NewUserController().ModifyUserInfoHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		}
	})
	
	//修改用户头像上传照片
	authorized.POST("/User/uploadPic", func(c *gin.Context){
		err := controller.NewUserController().UploadUserPicHandler(c)
		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		}
	})

	//修改管理员信息
	r.POST("/adminUser/modifyInfo", func(c *gin.Context) {
		err := controller.NewAdminUserController().ModifyAdminUserInfoHandler(c)
		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		}
	})

	//添加管理员信息
	r.POST("/adminUser/addInfo", func(c *gin.Context) {
		err := controller.NewAdminUserController().AddAdminUserHandler(c)
		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		}
	})


	//用户手机号注册
	r.POST("/register", func(c *gin.Context) {
		userID, token, err := controller.NewUserController().UserRegisterHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   300,
				"userId": nil,
				"token":  nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"userId": userID,
				"token":  token,
			})
		}
	})

	//管理员登录（未修改）
	/*r.POST("/admin/login", func(c *gin.Context) {
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
	})*/

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
		users, totalrecord := controller.NewUserController().AllUserInfoHandler(c)
		c.JSON(http.StatusOK, gin.H{
			"code":      200,
			"data":      users,
			"totals": totalrecord,
		},
		)
	})

	//获得特定名称用户信息
	r.GET("/User/theInfo", func(c *gin.Context) {
		users, err := controller.NewUserController().UserInfoHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"data": nil,
			},
			)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": users,
			},
			)
		}
	})

	//删除用户信息
	r.GET("/User/deleteInfo", func(c *gin.Context) {
		err := controller.NewUserController().DeleteUserInfoHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		}
	})

	//获得所有管理员信息
	r.GET("/adminUser/allInfo", func(c *gin.Context) {
		adminusers := controller.NewAdminUserController().AllAdminInfoHandler()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": adminusers,
		})
	})

	//获得特定管理员信息
	r.GET("/adminUser/theInfo", func(c *gin.Context) {
		adminuser,err := controller.NewAdminUserController().AdminInfoHandler(c)
		if err != nil{
			c.JSON(http.StatusOK, struct{
				*model.AdminUserInfo
				Code int `json:"code"`
			}{
				adminuser,
				300,
			})
		} else {
			c.JSON(http.StatusOK, struct{
				*model.AdminUserInfo
				Code int `json:"code"`
			}{
				adminuser,
				200,
			})
		}
	})

	//删除管理员信息
	r.GET("/adminUser/deleteInfo", func(c *gin.Context) {
		err := controller.NewAdminUserController().DeleteAdminUserInfoHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		}
	})


	//用户手机号登录
	r.GET("/login/cellphone", func(c *gin.Context) {
		userID, token, err := controller.NewUserController().UserLoginHandler(c)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"userId": userID,
				"token":  token,
			})
		} else if err.Error() == "手机号输入错误！" {
			c.JSON(http.StatusOK, gin.H{
				"code":   300,
				"userId": nil,
				"token":  nil,
			})
		} else if err.Error() == "密码输入错误！" {
			c.JSON(http.StatusOK, gin.H{
				"code":   301,
				"userId": nil,
				"token":  nil,
			})
		}
	})

	//获取登录用户信息
	authorized.GET("/user/profile", func(c *gin.Context) {
		userprofile, err := controller.NewUserController().UserProfileHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    300,
				"profile": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"profile": userprofile,
			})
		}
	})

	r.GET("/song/lyric", func(c *gin.Context) {
		lyric, err := controller.NewSongController().GetSongLyricHandler(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "没有此歌曲"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "获取歌词成功", "lyric": lyric})
		}
	})

	/*r.GET("/song/detail", func(c *gin.Context) {
		songname, singer, err := controller.NewSongController().GetSongDetailHandler(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "没有此歌曲"})
		} else {

			c.JSON(http.StatusOK, gin.H{"message": "获取信息成功", "songname": songname, "singer": singer})
		}
	})*/

	r.GET("/search/hot", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"searchHot": []gin.H{
				{
					"searchWord": "动物世界",
					"type":       1,
				},
				{
					"searchWord": "渡 The Crossing",
					"type":       3,
				},
				{
					"searchWord": "薛之谦",
					"type":       2,
				},
				{
					"searchWord": "上原れな",
					"type":       2,
				},
				{
					"searchWord": "Fill You",
					"type":       1,
				},
			},
		})

	})

	r.GET("/search/suggest", func(c *gin.Context) {
		albums, artists, songs := controller.NewSearchController().SearchSuggest(c)
		if albums == nil && artists == nil && songs == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    300,
				"albums":  nil,
				"artists": nil,
				"songs":   nil,
			})
		} else {
			messages_album := make([]gin.H, 0)
			for i := 0; i < len(albums) && i < 4; i++ {
				album := albums[i]
				message := gin.H{
					"albumid":      album.AlbumID,
					"album":        album.Name,
					"picUrl":       album.Pic_url,
					"artist":       album.Artist,
					"publish_time": album.Publish_time,
					"size":         album.Size,
					"type":         album.Type,
					"artist_id":    album.Artist_ID,
					"profile":      album.Profile,
				}
				messages_album = append(messages_album, message)
			}
			messages_artist := make([]gin.H, 0)
			for i := 0; i < len(artists) && i < 4; i++ {
				artist := artists[i]
				message := gin.H{
					"artistId":  artist.ArtistID,
					"artist":    artist.Name,
					"picUrl":    artist.Pic_url,
					"songSize":  artist.Song_size,
					"albumSize": artist.Album_size,
					"profile":   artist.Profile,
					"location":  artist.Location,
				}
				messages_artist = append(messages_artist, message)
			}
			message_song := make([]gin.H, 0)
			for i := 0; i < len(songs) && i < 4; i++ {
				song := songs[i]
				message := gin.H{
					"songId":      song.Song_ID,
					"name":        song.Name,
					"artist":      song.Artist,
					"album":       song.Album,
					"duration":    song.Duration,
					"albumId":     song.Album_ID,
					"artistId":    song.Artist_ID,
					"pop":         song.Pop,
					"mark":        song.Mark,
					"publishTime": song.Publish_time,
					"url":         song.Url,
					"lyricUrl":    song.Lyric_url,
					"picUrl":      song.Pic_url,
					"type":        song.Type,
				}
				message_song = append(message_song, message)
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"albums":  messages_album,
				"artists": messages_artist,
				"songs":   message_song,
			})
		}

	})
}
