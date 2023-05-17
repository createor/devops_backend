package test

import (
	"server/utils"
	"testing"
)

func TestSessionDial(t *testing.T) {
	test := utils.RemoteHost{
		Host:     "127.0.0.1",
		Port:     "22",
		UserName: "root",
		Password: "Txk@110#",
	}
	s, e := test.NewSSHSession("123456")
	if e != nil {
		t.Errorf("创建会话失败:%v", e)
	}
	out, err := utils.SessionDial(s, "whoami")
	if err != "" {
		t.Errorf("执行命令错误:%v", err)
	}
	if out != "root" {
		t.Errorf("获取结果错误:%v", out)
	}
}
