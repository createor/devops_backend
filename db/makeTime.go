package db

import (
	"time"
)

// 以固定格式返回当前时间
func NewTableTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
