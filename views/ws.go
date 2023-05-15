package views

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// 返回ws接口
func NewWebsocketHandler(fn1 func(*gin.Context), fn2 func(...string) string, flag bool) func(*gin.Context) {
	return func(c *gin.Context) {
		if fn1 != nil {
			fn1(c)
		}
		var host, name, data string
		if flag {
			host = c.Query("host")
			name = c.Query("name")
		}
		// 升级连接
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("升级协议失败:", err)
		}
		defer conn.Close()
		for {
			// 读取客户端发送的信息
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("读取客户端信息失败:", err)
				break
			}
			fmt.Println(string(msg))
			if flag {
				data = fn2(host, name, string(msg))
			} else {
				data = fn2(string(msg))
			}
			// 向客户端发送信息
			err = conn.WriteMessage(websocket.TextMessage, []byte(data))
			if err != nil {
				fmt.Println("发送信息失败:", err)
				break
			}
		}
	}
}
