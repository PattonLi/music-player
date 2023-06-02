package router

import (
	"errors"
	"fmt"
	"music-player/musicplayerserver/controller"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 定义JWT密钥
var jwtKey = []byte("1234567890")

// token载荷
type Claims struct {
	Username string `json:"username"`
	Admin    string `json:"admin"`
	jwt.StandardClaims
}

// token过期时间
var expireTime = time.Now().Add(time.Minute * 30)

// 生成token
func createToken(username string, admin string) (string, error) {
	claims := &Claims{
		Username: username,
		Admin:    admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "test.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		err = errors.New("token生成失败！")
	}
	return tokenString, err
}

// 鉴权中间件
func authMiddleware(c *gin.Context) {
	// 获取JWT Token
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	}
	// 验证Token是否有效
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Token not found"})
		c.Abort()
		return
	}
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization Token"})
		c.Abort()
		return
	}
	// 判断Token是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		c.Set("username", claims.Username)
		c.Set("admin", claims.Admin)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization Token2"})
		c.Abort()
		return
	}
}

func Posts(r *gin.Engine) {

	//鉴权路由，需要鉴权的post api将r改为authorized
	authorized := r.Group("/")
	authorized.Use(authMiddleware)
	//添加用户
	r.POST("/User/addInfo", func(c *gin.Context) {
		user, err := controller.NewUserController().AddUserHandler(c)
		var code int
		if err != nil {
			fmt.Print(err.Error())
			code = 300
		} else {
			code = 200
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
			tokenString, err = createToken(username, "false")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "token": tokenString, "adminid": userID})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "登录成功！", "token": tokenString, "userid": userID})
			}
		}
	})

	//管理员登录（未修改）
	r.POST("/admin/login", func(c *gin.Context) {
		adminname, adminID, err := controller.NewUserController().AdminLoginHandler(c)
		tokenString := ""
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "token": tokenString, "adminid": adminID})
		} else {
			tokenString, err = createToken(adminname, "true")
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
	authorized.Use(authMiddleware)

	//获得所有用户信息
	r.GET("/User/allInfo", func(c *gin.Context) {
		users := controller.NewUserController().AllUserInfoHandler()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": []gin.H{
				{
					"users": users,
				},
			},
		})
	})

	//获得特定用户信息
	r.GET("/User/theInfo", func(c *gin.Context) {

	})

	//删除用户信息
	r.GET("/User/deleteInfo", func(c *gin.Context) {

	})

	//获得所有管理员信息
	r.GET("/adminUser/allInfo", func(c *gin.Context) {
		adminusers := controller.NewUserController().AllAdminInfoHandler()
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
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
