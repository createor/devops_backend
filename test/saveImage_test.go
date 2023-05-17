package test

import (
	"io/ioutil"
	"server/utils"
	"testing"
)

func TestImgName(t *testing.T) {
	testStr := "test.png"
	result := utils.ImgName(testStr)
	if result == "" {
		t.Error("获取图片名称失败")
	}
}

func TestSaveImg(t *testing.T) {
	testImg := "test.png"
	testStr, _ := ioutil.ReadFile(testImg)
	testName := "success.png"
	result, _ := utils.SaveImg(testStr, testName)
	if result != true {
		t.Error("图片保存失败")
	}
}
