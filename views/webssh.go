package views

import (
	"fmt"
	"net/http"

	"server/db"
	"server/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

/*
ssh登陆主机转发
*/

// 连接前检查用户权限
func PreloadConnect(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		c.Abort()
		return
	}
	// 查询是否拥有主机权限
	host := c.Query("host")
	name := c.Query("name")
	if host == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		c.Abort()
		return
	}
	isExist, _ := db.QueryHostByName(host, name, userID.(string))
	if !isExist {
		c.JSON(http.StatusForbidden, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "无权限操作",
		})
		c.Abort()
		return
	}
}

// ssh连接主机
type WsSshData struct {
	ID      string
	Command string
}

func ConnectSSH(args ...string) string {
	var wd WsSshData
	var e error
	_ = json.Unmarshal([]byte(args[2]), &wd)
	if wd.ID == "" {
		wd.ID = utils.NewUuid()
	}
	s, ok := utils.SSHSessionPool[wd.ID]
	if !ok {
		// 查找端口和密码
		hostInfo, _ := db.QueryHost()
		rh := utils.RemoteHost{
			Host:     args[0],
			Port:     hostInfo.Port,
			UserName: args[1],
			Password: hostInfo.Password,
		}
		s, e = rh.NewSSHSession(wd.ID)
		if e != nil {
			return fmt.Sprintf("{'status':-1,'data':'','msg':'错误:%v'}", e)
		}
	}
	out, err := utils.SessionDial(s, args[2])
	if err != "" {
		return fmt.Sprintf("{'status':-1,'data':'','msg':'错误:%v'}", err)
	}
	return fmt.Sprintf("{'status':0,'data':'%v','msg':'成功'}", out)
}
