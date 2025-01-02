package gg

import (
	"fmt"
	"testing"
)

func TestByte01(t *testing.T) {
	// 示例：将 int 转换为 byte
	intValue := 128
	byteValue := Byte_intToByte(intValue)
	fmt.Printf("Int value: %d\n", intValue)
	fmt.Printf("Byte value: %d\n", byteValue)

	// 示例：超出 byte 范围的情况（触发 panic）
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	intValue = 300
	byteValue = Byte_intToByte(intValue)
	fmt.Printf("Int value: %d\n", intValue)
	fmt.Printf("Byte value: %d\n", byteValue)
}
