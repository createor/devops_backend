package test

import (
	"server/utils"
	"testing"
)

func TestNewCache(t *testing.T) {
	instance1 := utils.NewCache()
	instance2 := utils.NewCache()
	if instance1 != instance2 {
		t.Error("实例对象不一致")
	}
}
