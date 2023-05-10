package utils

import (
	"time"

	"sync"

	"github.com/patrickmn/go-cache"
)

/*
使用单例模式返回cache对象
*/

var (
	instance *cache.Cache
	once     sync.Once
)

func NewCache() *cache.Cache {
	once.Do(func() {
		// 每5分钟清理过期缓存，每10分钟清理过期项目
		instance = cache.New(5*time.Minute, 10*time.Minute)
	})
	return instance
}
