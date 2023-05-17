package views

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"server/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

/*
中间件
*/

// 密钥
var jwtKey = []byte(utils.Settings.Secert)

// 白名单路径
var ignorePath = utils.Settings.IgnorePath

// 缓存
var redis = utils.NewCache()

// 需要加密的信息
type Claims struct {
	UserID       string `json:"user_id"`       // 用户id
	UserName     string `json:"user_name"`     // 用户名,未加密形式
	DepartmentID string `json:"department_id"` // 部门id
	Department   string `json:"department"`    // 部门名
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
func createToken(userID, userName, departmentID, department string) (string, error) {
	// 先从缓存中查找
	if storeToken, found := redis.Get(userID); found {
		if newToken := storeToken.(string); newToken != "0" {
			return newToken, nil
		}
	}
	claims := &Claims{
		UserID:       userID,
		UserName:     userName,
		DepartmentID: departmentID,
		Department:   department,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(utils.Settings.Expire)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	// 缓存存储token
	redis.Set(userID, newToken, time.Duration(utils.Settings.Expire)*time.Hour)
	return newToken, nil
}

// 设置token过期
func ExpireToken(userID string) {
	redis.Set(userID, "0", -1)
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
					"msg":    "认证信息不能为空",
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
				if e, ok := err.(*jwt.ValidationError); ok {
					if e.Errors&jwt.ValidationErrorExpired != 0 {
						ctx.JSON(http.StatusUnauthorized, gin.H{
							"status": -1,
							"data":   "",
							"msg":    "token已过期",
						})
						ctx.Abort()
						return
					}
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"status": -1,
						"data":   "",
						"msg":    "认证失败",
					})
					ctx.Abort()
					return
				}
			}
			// 将用户信息添加到上下文中
			if claims, ok := token.Claims.(*Claims); ok && token.Valid {
				if storeToken, found := redis.Get(claims.UserID); found {
					if newToken := storeToken.(string); newToken == "0" {
						ctx.JSON(http.StatusUnauthorized, gin.H{
							"status": -1,
							"data":   "",
							"msg":    "token已失效",
						})
						ctx.Abort()
						return
					}
				}
				ctx.Set("user_id", claims.UserID)
				ctx.Set("user_name", claims.UserName)
				ctx.Set("department_id", claims.DepartmentID)
				ctx.Set("department", claims.Department)
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": -1,
					"data":   "",
					"msg":    "认证失败",
				})
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}

// 日志中间件
func logMiddleware() (*rotatelogs.RotateLogs, error) {
	logFilePath := utils.Settings.Logger.Path
	// 如果日志目录不存在则创建
	logFileName := filepath.Join(logFilePath, utils.Settings.Logger.Name)
	return rotatelogs.New(
		logFileName+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logFileName),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(uint(utils.Settings.Logger.Number)),
	)
}
