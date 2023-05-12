package views

import (
	"net/http"

	"server/db"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 创建部门
func AddDepartment(c *gin.Context) {
	db.CreateDepartment(db.Department{
		Uuid: utils.NewUuid(),
	})
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "创建成功",
	})
}

func GetAllUserFromDepartment(c *gin.Context) {
	// 判断用户是否是此部门或者此部门的上一级部门
	// 如果不是，返回403权限禁止
	departmentID := c.Query("id")
	if departmentID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	data, err := db.QueryAllUserByDepartment(departmentID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "查询失败",
		})
		return
	}
	if len(data) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"data":   "",
			"msg":    "查询为空",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   data,
		"msg":    "查询成功",
	})
}

// 用户绑定部门
func BindDepartment(c *gin.Context) {
	// 判断用户是否已经有部门
	// 如果用户没有绑定则可以插入，如果有，查看用户原部门是否被用户管理，不是则不能更新
}
