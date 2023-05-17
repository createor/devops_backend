package test

import (
	"server/utils"
	"testing"
)

func TestDecryptData(t *testing.T) {
	testStr := "CzeS9TtHbLJj8oBoZNgI5MzyIsbbzI2WIOOTBHBgv3TonXk9utn33AdNBOkvKh3AF5CVJ+52eM4QosSw6yir5H9Ueug5C5BMvUBKwIesiFP3+1v7lPHWZel7aDqNj+7/JI5TG3zplTl0BwnDvEhAo+QiXMGc/ThRgSme1E4jOiQONz9jjIEWNXMM5/f7BsMmnQeGRa7+LLdUOQLrWkqLOKsEyQufKp3yTHDHU02PTuL0gqqrP15Zt74UmZwWuxhV88UwzUlShW+YDKpXgIF5zJ1U+XEUMXlFdC5MSMVrzr/i3J/aausixX0ZkXuI2vRUcnm5g+G8mkuS1/c+3gSQYA=="
	want := "21232f297a57a5a743894a0e4a801fc3"
	result, err := utils.DecryptData(testStr)
	if result != want {
		t.Error("解密失败", err)
	}
}
