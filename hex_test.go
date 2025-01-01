package main

import (
	"fmt"
	"testing"
)

func TestHex(t *testing.T) {
	tests := []string{
		"0x123abc",
		"0XABCDEF",
		"123abc",
		"ABCDEF",
		"0x123g", // 包含非法字符
		"123g",   // 包含非法字符
		"",       // 空字符串
		"0x",     // 只有前缀
	}

	for _, test := range tests {
		fmt.Printf("'%s' --> %v\n", test, IsHexNumber(test))
	}

	// 示例十六进制字符数组
	hexArray := []string{"48", "65", "6C", "6C", "6F", "20", "57", "6F", "72", "6C", "64"}

	// 转换为字符串
	result, err := HexArrayToString(hexArray)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	// 输出结果
	fmt.Println("十六进制字符数组:", hexArray)
	fmt.Println("转换后的字符串:", result)
}
