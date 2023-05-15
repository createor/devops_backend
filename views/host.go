package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/db"
)

func GetAllHostBySelf(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	data, err := db.QueryAllHostByUser(userID.(string))
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
			"status": -1,
			"data":   "",
			"msg":    "查询为空",
		})
		return
	}
	result := db.MergeHost(data)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   result,
		"msg":    "查询成功",
	})
}
