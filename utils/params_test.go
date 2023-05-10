package utils

import (
	"testing"
)

func TestYaml2Struct(t *testing.T) {
	testFile := "F:/代码/devops_backend/server/server.yaml"
	_, err := Yaml2Struct(testFile)
	if err != nil {
		t.Errorf("读取yaml配置文件失败: %v", err)
	}
}
