package test

import (
	"encoding/json"
	"server/utils"
	"testing"
)

func TestCheck(t *testing.T) {
	testJson_1 := `{"address": "123456"}`
	testJson_2 := `{"name": "admin", "address": "123456"}`
	type TestStr struct {
		Name    string `json:"name" bind:"required"`
		Address string `json:"address"`
	}
	var want1 TestStr
	var want2 TestStr
	_ = json.Unmarshal([]byte(testJson_1), &want1)
	_ = json.Unmarshal([]byte(testJson_2), &want2)
	result_1 := utils.Check(want1)
	if result_1 {
		t.Errorf("校验json转换失败: %v", want1)
	}
	result_2 := utils.Check(want2)
	if !result_2 {
		t.Errorf("校验json转换失败: %v", want2)
	}
}
