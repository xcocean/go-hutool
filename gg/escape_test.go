package gg

import (
	"fmt"
	"testing"
)

func TestEscape01(t *testing.T) {
	// 示例：反转义 XML 中的特殊字符
	input := `&lt;message&gt;Hello, &quot;World&quot; &amp; &apos;Universe&apos;!&lt;/message&gt;`
	unescaped := Escape_UnescapeXML(input)
	fmt.Printf("Escaped: %s\n", input)
	fmt.Printf("Unescaped: %s\n", unescaped)
}

func TestEscape02(t *testing.T) {
	// 示例：转义 XML 中的特殊字符
	input := `<message>Hello, "World" & 'Universe'!</message>`
	escaped := Escape_escapeXML(input)
	fmt.Printf("Original: %s\n", input)
	fmt.Printf("Escaped: %s\n", escaped)
}

func TestEscape03(t *testing.T) {
	// 示例：转义 HTML4 中的特殊字符
	input := `<p>Hello, "World" & 'Universe'!</p>`
	escaped := Escape_escapeHTML4(input)
	fmt.Printf("Original: %s\n", input)
	fmt.Printf("Escaped: %s\n", escaped)
}

func TestEscape04(t *testing.T) {
	// 示例：反转义 HTML4 中的特殊字符
	input := `&lt;p&gt;Hello, &quot;World&quot; &amp; &apos;Universe&apos;!&lt;/p&gt;`
	unescaped := Escape_UnescapeHTML4(input)
	fmt.Printf("Escaped: %s\n", input)
	fmt.Printf("Unescaped: %s\n", unescaped)
}

func TestEscape05(t *testing.T) {
	// 示例：Unicode 转义编码
	input := "Hello, 世界！"
	escaped := Escape_escapeUnicode(input)
	fmt.Printf("Original: %s\n", input)
	fmt.Printf("Escaped: %s\n", escaped)
}

func TestEscape06(t *testing.T) {
	// 示例：Unicode 反转义解码
	input := `"Hello, \u4e16\u754c\uff01"`
	unescaped, err := Escape_UnescapeUnicode(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Escaped: %s\n", input)
		fmt.Printf("Unescaped: %s\n", unescaped)
	}
}
