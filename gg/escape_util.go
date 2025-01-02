package gg

import (
	"strconv"
	"strings"
)

// 转义 XML 中的特殊字符
func Escape_escapeXML(input string) string {
	// 定义需要转义的字符及其对应的转义序列
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&apos;",
	}
	// 逐个替换特殊字符
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}
	return input
}

// 反转义 XML 中的特殊字符
func Escape_UnescapeXML(input string) string {
	// 定义转义序列及其对应的原始字符
	replacements := map[string]string{
		"&amp;":  "&",
		"&lt;":   "<",
		"&gt;":   ">",
		"&quot;": "\"",
		"&apos;": "'",
	}

	// 逐个替换转义序列
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}
	return input
}

// 转义 HTML4 中的特殊字符
func Escape_escapeHTML4(input string) string {
	// 定义需要转义的字符及其对应的转义序列
	replacements := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&apos;",
	}

	// 逐个替换特殊字符
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}

	return input
}

// 反转义 HTML4 中的特殊字符
func Escape_UnescapeHTML4(input string) string {
	// 定义转义序列及其对应的原始字符
	replacements := map[string]string{
		"&amp;":  "&",
		"&lt;":   "<",
		"&gt;":   ">",
		"&quot;": "\"",
		"&apos;": "'",
	}

	// 逐个替换转义序列
	for old, new := range replacements {
		input = strings.ReplaceAll(input, old, new)
	}

	return input
}

// 对字符串进行 Unicode 转义编码
// Hello, 世界！ --> Hello, \u4e16\u754c\uff01
func Escape_escapeUnicode(input string) string {
	return strconv.QuoteToASCII(input)
}

// 对字符串进行 Unicode 反转义解码
//
//	input := `"Hello, \u4e16\u754c\uff01"`
//	unescaped, err := Escape_UnescapeUnicode(input)
func Escape_UnescapeUnicode(input string) (string, error) {
	// 如果输入字符串包含双引号，去除它们
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		input = input[1 : len(input)-1]
	}

	// 逐个字符解码
	var result strings.Builder
	for len(input) > 0 {
		r, multibyte, tail, err := strconv.UnquoteChar(input, '"')
		if err != nil {
			return "", err
		}
		result.WriteRune(r)
		input = tail
		if multibyte {
			// 如果是多字节字符，跳过已处理的部分
			continue
		}
	}
	return result.String(), nil
}
