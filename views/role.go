package views

import (
	"net/http"
	"server/db"
	"server/utils"

	"github.com/gin-gonic/gin"
)

/*
用户只能授权自己绑定的角色和菜单
*/

// 菜单路径
type Menu struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path,omitempty"`
	Icon     string `json:"icon"`
	Children []Menu `json:"children,omitempty"`
}

//
type NewRole struct {
	RoleID   string `json:"role_id" bind:"required"`
	RoleName string `json:"role_name" bind:"required"`
}

// 将角色绑定的菜单存储到内存
// go-cache

// 新增角色视图
func AddRole(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	var r NewRole
	err := c.BindJSON(&r)
	if err != nil || !utils.Check(r) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	createUserId := userID.(string)
	isExist, err := db.CreateRole(createUserId, db.Role{
		Uuid:     utils.NewUuid(),
		RoleID:   r.RoleID,
		RoleName: r.RoleName,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "创建角色失败",
		})
		return
	}
	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"data":   "",
			"msg":    "角色已存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "创建角色成功",
	})
}

// 绑定角色
type Role2User struct {
	RoleID string `json:"role_id" bind:"required"`
	UserID string `json:"user_id" bind:"required"`
}

func BindRole(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	var ru Role2User
	err := c.BindJSON(&ru)
	if err != nil || !utils.Check(ru) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	// 获取用户id和角色id
	newUserID := userID.(string)
	isExist, err := db.CreateRoleWithUser(newUserID, db.RoleWithUser{
		RoleID: ru.RoleID,
		UserID: ru.UserID,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "绑定角色失败",
		})
		return
	}
	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"data":   "",
			"msg":    "请勿重复绑定同一角色",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "绑定角色成功",
	})
}

// 获取用户绑定的角色
func GetUserRole(c *gin.Context) {
	reqUserId := c.Query("user_id")
	if reqUserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	data, isExist, err := db.QueryRoleByUser(reqUserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "请求失败",
		})
		return
	}
	if !isExist || len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "未查询到用户角色",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   data,
		"msg":    "请求成功",
	})
}
