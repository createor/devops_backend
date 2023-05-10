package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 图片保存目录
var ImgDirectory = "F:/代码/devops_backend/server"

// ImgName 获取图片名称，名称使用随机码
//
// 参数:
//
//	name: string, 图片原始名称
//
// 返回:
//
//	string: 随机图片名称
func ImgName(name string) string {
	// 获取图片后缀
	fileExt := filepath.Ext(name)
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)
	newFileName := hex.EncodeToString(randomBytes)
	return newFileName + fileExt
}

// SaveImg 保存图片到指定目录
func SaveImg(image []byte, name string) (bool, error) {
	fileName := filepath.Join(ImgDirectory, name)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err = ioutil.WriteFile(fileName, image, 0644)
		if err != nil {
			return false, err
		}
	} else {
		return false, fmt.Errorf("文件已存在")
	}
	return true, nil
}
