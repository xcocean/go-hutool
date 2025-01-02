package gg

import (
	"reflect"
	"strconv"
)

// 判断值是否为数字类型
func Number_isNumber[T any](obj T) bool {
	// 获取值的类型
	kind := reflect.TypeOf(obj).Kind()

	// 判断是否为数字类型
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// 判断值是否为数字类型，支持自定义数字类型
func Number_isNumber2[T any](obj T) bool {
	// 获取值的类型
	kind := reflect.TypeOf(obj).Kind()

	// 判断是否为数字类型
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		// 检查是否为自定义数字类型
		value := reflect.ValueOf(obj)
		if value.CanInt() || value.CanUint() || value.CanFloat() {
			return true
		}
		return false
	}
}

// 判断字符串是否为数字
func Number_isNumberByString(str string) bool {
	// 尝试将字符串解析为浮点数
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
