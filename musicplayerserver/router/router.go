package router

import (
	"fmt"
	"music-player/musicplayerserver/controller"
	"music-player/musicplayerserver/utils/jwt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Posts(r *gin.Engine) {

	//鉴权路由，需要鉴权的post api将r改为authorized
	authorized := r.Group("/")
	authorized.Use(utils.AuthMiddleware)
	//添加用户
	r.POST("/User/addInfo", func(c *gin.Context) {
		totals, currentPage, userlist, err := controller.NewUserController().AddUserHandler(c)
		var code int
		if err != nil {
			fmt.Print(err.Error())
			code = 300
		} else {
			code = 200
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"totals":totals,
			"currentPage": currentPage,
			"data": userlist,
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

	//管理员修改用户头像
	r.POST("/User/uploadPic", func(c *gin.Context){
		url, err := controller.NewUserController().AdminUploadUserPicHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"picUrl":url,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"picUrl":url,
			})
		}
	})

	//用户自行修改头像
	authorized.POST("/user/profile/edit", func(c *gin.Context){
		url, err := controller.NewUserController().UploadUserPicHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"picUrl":url,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"picUrl": url,
			})
		}
	})
	//修改管理员信息
	r.POST("/adminUser/modifyInfo", func(c *gin.Context) {
		err := controller.NewAdminUserController().ModifyAdminUserInfoHandler(c)
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

	//添加管理员信息
	r.POST("/adminUser/addInfo", func(c *gin.Context) {
		totals, currentPage, adminlist, err := controller.NewAdminUserController().AddAdminUserHandler(c)
		var code int
		if err != nil {
			code = 300
		} else {
			code = 200
		}
		c.JSON(200,gin.H{
			"code": code,
			"totals": totals,
			"currentPage": currentPage,
			"data": adminlist,
		})
	})

	//用户手机号注册
	r.POST("/register", func(c *gin.Context) {
		userID, token, err := controller.NewUserController().UserRegisterHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   300,
				"userId": 0,
				"token":  "",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"userId": userID,
				"token":  token,
			})
		}
	})



	//添加歌曲
	r.POST("/admin/addSong", func(c *gin.Context) {
		result := controller.NewSongController().AddSongHandler(c)
		if result {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 300,
			})
		}
	})

	//发送播放歌曲日志
	r.POST("/log/play", func(c *gin.Context) {
		err := controller.NewLogController().AddPlayLogHandler(c)
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

	//发送下载歌曲日志
	r.POST("/log/download", func(c *gin.Context) {
		err := controller.NewLogController().AddDownloadLog(c)
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

	//添加用户收藏
	r.POST("/user/like", func(c *gin.Context) {
		err := controller.NewLikesController().AddUserLikesHandler(c)
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

	//删除用户收藏
	r.POST("/user/like/delete", func(c *gin.Context) {
		err := controller.NewLikesController().DeleteUserLikesHandler(c)
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

	//发表评论
	r.POST("/song/comment/publish", func(c *gin.Context) {
		err := controller.NewCommentController().PublishCommentHandler(c)
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
				"data": users,
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

	//获得表中第index个用户信息
	r.GET("/User/aInfo", func(c *gin.Context) {
		totals, user, err := controller.NewUserController().GetAInfoHandler(c)
		code := 0;
		if(err != nil){
			code = 300
		} else {
			code = 200
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"totals": totals,
			"data": user,
		})

	})

	//获得所有管理员信息
	r.GET("/adminUser/pageAllInfo", func(c *gin.Context) {
		totals, adminusers := controller.NewAdminUserController().AllAdminInfoHandler(c)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"totals": totals,
			"data": adminusers,
		})
	})

	//获得特定管理员信息
	r.GET("/adminUser/theInfo", func(c *gin.Context) {
		adminlist, err := controller.NewAdminUserController().AdminInfoHandler(c)
		var code int
		if err != nil {
			code = 300
		} else {
			code = 200
		}
		c.JSON(200, gin.H{
			"code":code,
			"data": adminlist,
		})
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
	r.GET("/user/profile", func(c *gin.Context) { //authorized
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

	//管理员登录
	r.GET("/adminUser/login", func(c *gin.Context) {
		adminID, token, err := controller.NewAdminUserController().AdminLoginHandler(c)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"adminId": adminID,
				"token":  token,
			})
		} else if err.Error() == "账号输入错误！" {
			c.JSON(http.StatusOK, gin.H{
				"code":   300,
				"adminId": 0,
				"token":  "",
			})
		} else if err.Error() == "密码输入错误！" {
			c.JSON(http.StatusOK, gin.H{
				"code":   301,
				"adminId": 0,
				"token":  "",
			})
		}
	})

	//获取单个管理员信息
	r.GET("/adminUser/profile", func(c *gin.Context) {
		adminname, err := controller.NewAdminUserController().AdminProfileHandler(c)
		var code int
		if err != nil{
			code = 300
		} else {
			code = 200
		}
		c.JSON(200, gin.H{
			"code": code,
			"adminName": adminname,
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

	/*r.GET("/song/detail", func(c *gin.Context) {
		songname, singer, err := controller.NewSongController().GetSongDetailHandler(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "没有此歌曲"})
		} else {

			c.JSON(http.StatusOK, gin.H{"message": "获取信息成功", "songname": songname, "singer": singer})
		}
	})*/

	//未输入时显示热搜
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

	//搜索推荐
	r.GET("/search/suggest", func(c *gin.Context) {
		albums, artists, songs := controller.NewSearchController().SearchSuggestHandler(c)
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

	//歌曲推荐
	r.GET("/discover/song", func(c *gin.Context) {
		songs := controller.NewSongController().GetTenSongsHandler(c)
		if songs == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  300,
				"songs": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":  200,
				"songs": songs,
			})
		}
	})

	//专辑推荐
	r.GET("discover/album", func(c *gin.Context) {
		albums := controller.NewAlbumController().GetTenAlbumsHandler(c)
		if albums == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   300,
				"albums": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"albums": albums,
			})
		}
	})

	//歌手推荐
	r.GET("discover/artist", func(c *gin.Context) {
		artists := controller.NewArtistController().GetTenArtistHandler(c)
		if artists == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    300,
				"artists": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"artists": artists,
			})
		}
	})

	//专辑详细页信息
	r.GET("/detail/album", func(c *gin.Context) {
		songs, err0 := controller.NewSongController().GetAlbumSongsHandler(c)
		album, err1 := controller.NewAlbumController().GetAlbumByIdHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  300,
				"songs": nil,
				"album": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":  200,
				"songs": songs,
				"album": album,
			})
		}

	})

	//歌手详细页信息
	r.GET("/detail/artist", func(c *gin.Context) {
		artist, err := controller.NewArtistController().GetArtistDetailHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":         200,
				"artistDetali": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":         200,
				"artistDetail": artist,
			})
		}
	})
	//获得特定页所有专辑信息
	r.GET("/admin/pageAllAlbum", func(c *gin.Context) {
		albums, totalPage := controller.NewAlbumController().AllAlbumInfoHandler(c)
		c.JSON(http.StatusOK, gin.H{
			"code":      200,
			"data":      albums,
			"totalPage": totalPage,
		},
		)
	})

	//获得特定名称专辑信息
	r.GET("/admin/theAlbum", func(c *gin.Context) {
		albums, err := controller.NewAlbumController().AlbumInfoHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"data": nil,
			},
			)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": albums,
			},
			)
		}
	})

	//轮播图推荐
	r.GET("/discover/swiper", func(c *gin.Context) {
		type swiper struct {
			TargetId   int    `json:"targetId"`
			PicUrl     string `json:"picUrl"`
			TargetType int    `json:"targetType"`
			TypeTitle  string `json:"typeTitle"`
		}

		songs := controller.NewSongController().GetTenSongsHandler(c)
		albums := controller.NewAlbumController().GetTenAlbumsHandler(c)

		var swipers []swiper

		for i := 0; i < 3; i++ {
			var s1 swiper
			var s2 swiper

			s1.TargetId = songs[i].Song_ID
			s1.PicUrl = songs[i].Pic_url
			s1.TargetType, _ = strconv.Atoi(songs[i].Type)
			s1.TypeTitle = songs[i].Name

			s2.TargetId = albums[i].AlbumID
			s2.PicUrl = albums[i].Pic_url
			s2.TargetType, _ = strconv.Atoi(albums[i].Type)
			s2.TypeTitle = albums[i].Name

			swipers = append(swipers, s1)
			swipers = append(swipers, s2)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"swipers": swipers,
		})
	})

	//分页获取歌手歌曲
	r.GET("/detail/artist/songs", func(c *gin.Context) {
		songpage, err0, err1, pagetotal := controller.NewSongController().GetSongsPageHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"songs":     nil,
				"pageTotal": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"songs":     songpage,
				"pageTotal": pagetotal,
			})
		}
	})

	//分页获取歌手的专辑
	r.GET("/detail/artist/album", func(c *gin.Context) {
		albumpage, err0, err1, pagenum := controller.NewAlbumController().GetAlbumPageHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"albums":    nil,
				"pageTotal": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"albums":    albumpage,
				"pageTotal": pagenum,
			})
		}
	})

	//根据songid获取歌曲
	r.GET("/detail/song", func(c *gin.Context) {
		song, err := controller.NewSongController().GetSongDetailHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"song": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"song": song,
			})
		}
	})

	//获取歌手描述
	r.GET("/detail/artist/describe", func(c *gin.Context) {
		type Describe struct {
			Profile   string `json:"profile"`
			BasicInfo string `json:"basicInfo"`
			TimeLine  string `json:"timeLine"`
			Award     string `json:"award"`
		}

		profile, err := controller.NewArtistController().GetArtistDescribeHandler(c)

		if err != nil {
			var describe Describe
			describe.Profile = ""
			describe.BasicInfo = ""
			describe.TimeLine = ""
			describe.Award = ""

			c.JSON(http.StatusOK, gin.H{
				"code":     300,
				"describe": describe,
			})
		} else {
			var describe Describe
			describe.Profile = profile
			describe.BasicInfo = "test"
			describe.TimeLine = "test"
			describe.Award = "test"
			c.JSON(http.StatusOK, gin.H{
				"code":     200,
				"describe": describe,
			})
		}
	})

	//获取热门歌单
	r.GET("/library/playlist/hot", func(c *gin.Context) {
		playlist, err0, err1, pagatotal := controller.NewPlaylistController().GetHotPlaylistHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"pageTotal": nil,
				"platlist":  nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"pageTotal": pagatotal,
				"playlist":  playlist,
			})
		}
	})

	//获取歌单详情
	r.GET("/detail/playlist", func(c *gin.Context) {
		playlist, songs, err := controller.NewPlaylistController().GetPlaylistAllSongsHandler(c)
		if !err {
			c.JSON(http.StatusOK, gin.H{
				"code":     300,
				"playlist": nil,
				"songs":    nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":     200,
				"playlist": playlist,
				"songs":    songs,
			})
		}
	})

	// 关键词搜索歌曲
	r.GET("/search/song", func(c *gin.Context) {
		songpage, err0, err1, pagetotal := controller.NewSongController().GetSongByKeyWordHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"pageTotal": nil,
				"songs":     nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"pageTotal": pagetotal,
				"songs":     songpage,
			})
		}
	})

	//关键词搜索歌手
	r.GET("/search/artist", func(c *gin.Context) {
		artistpage, err0, err1, pagetotal := controller.NewArtistController().GetArtistByKeyWordHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"pageTotal": nil,
				"artists":   nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"pageTotal": pagetotal,
				"artists":   artistpage,
			})
		}
	})

	//关键词搜索专辑
	r.GET("/search/album", func(c *gin.Context) {
		albumpage, err0, err1, pagetotal := controller.NewAlbumController().GetAlbumByKeyWordHandler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"pageTotal": nil,
				"albums":    nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"pageTotal": pagetotal,
				"albums":    albumpage,
			})
		}
	})

	// 获取所有mv信息
	r.GET("/mv", func(c *gin.Context) {
		mvs, err := controller.NewMvController().GetAllMvHandler()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"mvs":  nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"mvs":  mvs,
			})
		}
	})

	// 获取用户的所有收藏
	r.GET("/user/like", func(c *gin.Context) {
		songs, albums, artists, playlists, err := controller.NewLikesController().GetUserLikesHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"songs":     nil,
				"albums":    nil,
				"artists":   nil,
				"playlists": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"songs":     songs,
				"albums":    albums,
				"artists":   artists,
				"playlists": playlists,
			})
		}
	})

	r.GET("/library/artist", func(c *gin.Context) {
		atrtists, err0, err1, pagetotal := controller.NewArtistController().GetAtistByCHanddler(c)
		if err0 != nil || err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":      300,
				"pageTotal": nil,
				"artists":   nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":      200,
				"pageTotal": pagetotal,
				"artists":   atrtists,
			})
		}
	})

	//获取特定歌曲评论
	r.GET("/comment/getComment", func(c *gin.Context) {
		frontendcommentlist, err := controller.NewCommentController().GetSongCommentHandler(c)
		var code int
		if err != nil{
			code = 300
		} else {
			code = 200
		}
		c.JSON(200, gin.H{
			"code": code,
			"data": frontendcommentlist,
		})
	})

	//删除特定评论
	r.GET("/comment/deleteComment", func(c *gin.Context) {
		err := controller.NewCommentController().DeltetCommentHandler(c)
		var code int
		if err != nil{
			code = 300
		} else {
			code = 200
		}
		c.JSON(200, gin.H{
			"code": code,
		})
	})

	//根据type获得日志
	r.GET("/admin/getLog",func(c *gin.Context) {
		loglist, err := controller.NewLogController().GetLogHandler(c)
		var code int
		if err != nil{
			code = 300
		} else {
			code = 200
		}
		c.JSON(200, gin.H{
			"code": code,
			"data": loglist,
		})
	})
}
