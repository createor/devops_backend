package test

import (
	"fmt"
	"server/utils"
	"testing"
)

func TestYaml2Struct(t *testing.T) {
	testFile := "/usr/local/share/vscode/devops/server/server.yaml"
	result, err := utils.Yaml2Struct(testFile)
	if err != nil {
		t.Errorf("读取yaml配置文件失败: %v", err)
	} else {
		fmt.Println(result)
	}
}
