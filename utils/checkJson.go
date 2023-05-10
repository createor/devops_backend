package utils

import (
	"reflect"
)

// Check 比较转换后的json和结构体是否一致
//
// 参数:
//
//	source: any, 实例化后的结构体
//
// 返回:
//
//	bool: true--有效,false--无效
func Check(source interface{}) bool {
	// 通过反射获取bind:required判断是否必须
	typeOfSource := reflect.TypeOf(source)
	valueOfSource := reflect.ValueOf(source)
	for i := 0; i < typeOfSource.NumField(); i++ {
		field := typeOfSource.Field(i)
		fieldTag := field.Tag.Get("bind") // 获取标签
		if fieldTag == "required" {       // 当标签为必要时
			fieldValue := valueOfSource.Field(i)                                // 获取值
			if fieldValue.Interface() == reflect.Zero(field.Type).Interface() { // 判断值是否为空
				return false
			}
		}
	}
	return true
}
