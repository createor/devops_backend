package utils

import (
	"testing"
)

func TestNewCache(t *testing.T) {
	instance1 := NewCache()
	instance2 := NewCache()
	if instance1 != instance2 {
		t.Error("实例对象不一致")
	}
}
