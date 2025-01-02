package gg

import "fmt"

// 将 int 转换为 byte，如果超出范围则 panic
func Byte_intToByte(intValue int) byte {
	// 检查 intValue 是否在 byte 的范围内
	if intValue < 0 || intValue > 255 {
		panic(fmt.Sprintf("int value %d is out of byte range (0-255)", intValue))
	}

	// 转换为 byte
	return byte(intValue)
}

// 将 byte 转换为 int
func Byte_byteToInt(byteValue byte) int {
	return int(byteValue)
}
