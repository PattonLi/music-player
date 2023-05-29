package router

import (
	"errors"
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

	r.POST("/register", func(c *gin.Context) {
		err := controller.NewUserController().AddUserHandler(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "注册成功！"})
		}
	})

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
