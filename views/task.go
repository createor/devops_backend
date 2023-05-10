package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 任务完成后的回调函数
// POST
// http://ip:port/api/task/result/xxxx,xxx为任务id
type TaskResult struct {
	Status string `json:"status"`           // 成功或者失败
	ErrMsg string `json:"errmsg,omitempty"` // 错误信息,如果失败的话
}

func TaskCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
	})
}
