package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// 生成32位uuid
func NewUuid() string {
	id := uuid.NewV4()
	return strings.Replace(id.String(), "-", "", -1) // 替换uuid中的-符号
}
