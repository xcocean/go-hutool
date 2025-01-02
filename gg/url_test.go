package gg

import (
	"fmt"
	"testing"
)

func TestURL01(t *testing.T) {
	// 示例：URL 编码和解码
	input := "https://example.com?name=张三"

	// URL 编码
	encoded := URL_Encode(input)
	fmt.Printf("Encoded: %s\n", encoded)

	// URL 解码
	decoded := URL_Decode(encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
