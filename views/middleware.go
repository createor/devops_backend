package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 密钥
var jwtKey = []byte("123456")

// 白名单路径
var ignorePath = []string{
	"/api/user/login",
	"/api/user/getkey",
}

type Claims struct { // 需要加密的信息
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// 比较路径是否存在于白名单中
func isIgnore(path string) bool {
	for _, s := range ignorePath {
		if s == path {
			return true
		}
	}
	return false
}

// 创建token
func createToken(userID, userName string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// jwt认证中间件
func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求路径
		reqPath := ctx.Request.URL.Path
		// 判断请求路径是否在
		if !isIgnore(reqPath) { // 如果路径不在白名单中进行校验
			// 获取auth头部信息
			authHeader := ctx.GetHeader("Authorization")
			if authHeader == "" { // 不存在或者为空
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"data":   "",
					"msg":    "认证失败",
				})
				ctx.Abort()
				return
			}
			// 解析token
			token, err := jwt.ParseWithClaims(authHeader, &Claims{}, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("未知加密算法: %v", token.Header["alg"])
				}
				return jwtKey, nil
			})
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": -1,
					"data":   "",
					"msg":    "认证失败",
				})
				ctx.Abort()
				return
			}
			// 将用户信息添加到上下文中
			if claims, ok := token.Claims.(*Claims); ok && token.Valid {
				ctx.Set("user_id", claims.UserID)
				ctx.Set("user_name", claims.UserName)
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": -1,
					"data":   "",
					"msg":    "认证失败",
				})
				return
			}
		}

		ctx.Next()
	}
}
