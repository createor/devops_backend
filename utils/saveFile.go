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

var FilePath = ""

func GetFileNameByRandom() string {
	randomBytes := make([]byte, 30)
	_, _ = rand.Read(randomBytes)
	return hex.EncodeToString(randomBytes)
}

func FilterArticle(source string) []byte {
	safeHTML := template.HTMLEscapeString(source)
	return []byte(safeHTML)
}

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
