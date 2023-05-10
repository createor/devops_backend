package utils

import (
	"io/ioutil"
	"testing"
)

func TestImgName(t *testing.T) {
	testStr := "test.png"
	result := ImgName(testStr)
	if result == "" {
		t.Error("获取图片名称失败")
	}
}

func TestSaveImg(t *testing.T) {
	testImg := "test.png"
	testStr, _ := ioutil.ReadFile(testImg)
	testName := "success.png"
	result, _ := SaveImg(testStr, testName)
	if result != true {
		t.Error("图片保存失败")
	}
}
