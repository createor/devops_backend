package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 文件保存路径
var FilePath = filepath.Join(Settings.FilePath, "files")

// 获取30位的随机码作为文章名
func GetFileNameByRandom() string {
	randomBytes := make([]byte, 30)
	_, _ = rand.Read(randomBytes)
	return hex.EncodeToString(randomBytes)
}

// 过滤文章内容中不安全部分
func FilterArticle(source string) []byte {
	safeHTML := template.HTMLEscapeString(source)
	return []byte(safeHTML)
}

// 保存文章内容
func SaveArticle(file []byte, name string) (bool, error) {
	fileName := filepath.Join(FilePath, name)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err = ioutil.WriteFile(fileName, file, 0644)
		if err != nil {
			return false, err
		}
	} else {
		return false, fmt.Errorf("文件已存在")
	}
	return true, nil
}

// 展示文章内容
func ShowArtcle(name string) string {
	source, _ := ioutil.ReadFile(filepath.Join(FilePath, name))
	return string(source)
}
