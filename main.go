package main

import (
	"fmt"
	"net/http"

	"server/utils"
	route "server/views"
)

/*
主文件、入口文件
*/

func main() {
	s := &http.Server{
		Addr:    ":" + fmt.Sprintf("%d", utils.Settings.Port), // 运行端口
		Handler: route.InitRouter(),
	}

	s.ListenAndServe()
}
