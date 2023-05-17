package test

import (
	"fmt"
	"server/utils"
	"testing"
)

func TestReadTable(t *testing.T) {
	testFile := "/usr/local/share/vscode/devops/server/test.xlsx"
	result, err := utils.ReadTable(testFile, 2)
	if err != nil {
		t.Errorf("读取表格失败:%v", err)
	} else {
		fmt.Println(result)
	}
}
