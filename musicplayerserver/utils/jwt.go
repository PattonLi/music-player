package utils

import (
	"errors"
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
func CreateToken(username string, admin string) (string, error) {
	claims := &Claims{
		Username: username,
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
func AuthMiddleware(c *gin.Context) {
	// 获取JWT Token
	tokenString := c.GetHeader("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	}
	// 验证Token是否存在
	if tokenString == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":0,
			"data":[]gin.H{},
		})
		c.Abort()
		return
	}
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":0,
			"data":[]gin.H{},
		})
		c.Abort()
		return
	}
	// 判断Token是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		c.Set("username", claims.Username)
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":0,
			"data":[]gin.H{},
		})
		c.Abort()
		return
	}
}