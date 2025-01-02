package gg

import (
	"fmt"
	"regexp"
	"strings"
)

// 字符串是否为空白
// IsBlank("") // true
// IsBlank(" \t\n") // true
// IsBlank("abc") // false
func Str_isBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 字符串是否为非空白
// IsNotBlank("") // false
// IsNotBlank("") // false
// IsNotBlank(" \t\n") // false
// IsNotBlank("abc") // true
func Str_isNotBlank(str string) bool {
	return !Str_isBlank(str)
}

// 字符串是否为空
// IsEmpty("") // true
// IsEmpty(" \t\n") // false
// IsEmpty("abc") // false
func Str_isEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 字符串是否为非空白
// IsNotEmpty("") // false
// IsNotEmpty(" \t\n") // true
// IsNotEmpty("abc") // true
func Str_isNotEmpty(str string) bool {
	return !Str_isEmpty(str)
}

// 是否包含空字符串
// 如果指定的字符串数组的长度为 0，或者其中的任意一个元素是空字符串，则返回 true。
func Str_hasEmpty(str ...string) bool {
	if len(str) == 0 {
		return true
	}
	for _, s := range str {
		if Str_isEmpty(s) {
			return true
		}
	}
	return false
}

// 指定字符串数组中的元素，是否全部为空字符串。
func Str_isAllEmpty(str ...string) bool {
	if len(str) == 0 {
		return true
	}
	for _, s := range str {
		if !Str_isEmpty(s) {
			return false
		}
	}
	return true
}

// 指定字符串数组中的元素，是否全部为空字符串。
func Str_isAllBlank(str ...string) bool {
	if len(str) == 0 {
		return true
	}
	for _, s := range str {
		if !Str_isBlank(s) {
			return false
		}
	}
	return true
}

// 除去字符串头尾部的空白
//
//	Trim("")            = ""
//	Trim("     ")       = ""
//	Trim("abc")         = "abc"
//	Trim("    abc    ") = "abc"
func Str_Trim(str string) string {
	return strings.TrimSpace(str)
}

/**
 * 除去字符串头部的空白，如果字符串是{@code ""}，则返回{@code ""}。
 */
func Str_trimStart(str string) string {
	return strings.TrimLeft(str, " ")
}

/**
 * 除去字符串尾部的空白，如果字符串是{@code ""}，则返回{@code ""}。
 */
func Str_trimEnd(str string) string {
	return strings.TrimRight(str, " ")
}

// 是否以指定字符串开头
func Str_startWith(str string, c string) bool {
	if Str_isEmpty(str) {
		return false
	}
	return strings.HasPrefix(str, c)
}

// 字符串是否以给定字符结尾
func Str_endWith(str string, c string) bool {
	if Str_isEmpty(str) {
		return false
	}
	return strings.HasSuffix(str, c)
}

// 是否以指定字符串开头，忽略大小写
func Str_startWithIgnoreCase(s, prefix string) bool {
	if Str_isEmpty(s) {
		return false
	}
	return Str_startWith(strings.ToLower(s), strings.ToLower(prefix))
}

// 是否以指定字符串结尾，忽略大小写
func EndWithIgnoreCase(s, suffix string) bool {
	if Str_isEmpty(s) {
		return false
	}
	return Str_endWith(strings.ToLower(s), strings.ToLower(suffix))
}

func Str_equals(s, c string) bool {
	return s == c
}

func Str_equalsIgnoreCase(s, c string) bool {
	return strings.ToLower(s) == strings.ToLower(c)
}

// 字节数组转UTF-8字符串
func Str_utf8Str(byt []byte) string {
	return string(byt)
}

func Str_toString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

// 删除字符串中的所有空白字符
func Str_removeSpace(str string) string {
	// 编译正则表达式，匹配所有空白字符
	re := regexp.MustCompile(`\s+`)
	// 将匹配到的空白字符替换为空字符串
	return re.ReplaceAllString(str, "")
}

// 替换字符串中指定的字符
func Str_replaceChar(input string, oldChar, newChar rune) string {
	// 将字符串转换为 rune 切片以便处理 Unicode 字符
	inputRunes := []rune(input)
	for i, ch := range inputRunes {
		if ch == oldChar {
			inputRunes[i] = newChar
		}
	}
	return string(inputRunes)
}

// 替换字符串中指定的子字符串
func Str_replaceString(input, oldStr, newStr string) string {
	return strings.ReplaceAll(input, oldStr, newStr)
}

// 字符串转大写
func Str_toUpper(str string) string {
	return strings.ToUpper(str)
}

// 字符串转小写
func Str_toLower(str string) string {
	return strings.ToLower(str)
}

// 查找子字符串在字符串中的位置，未找到返回 -1
func Str_indexOf(str, substr string) int {
	return strings.Index(str, substr)
}

// LastIndexOf 查找子字符串在字符串中的最后一个匹配位置，未找到返回 -1
func Str_lastIndexOf(str, substr string) int {
	return strings.LastIndex(str, substr)
}

// 从字符串中提取子字符串
func Str_substring(str string, start, end int) string {
	// 将字符串转换为 rune 切片以便正确处理 Unicode 字符
	runes := []rune(str)

	// 检查索引范围
	if start < 0 || end > len(runes) || start > end {
		panic("Invalid start or end index")
	}

	// 提取子字符串
	return string(runes[start:end])
}

func Str_substring2(str string, start int) string {
	// 将字符串转换为 rune 切片以便正确处理 Unicode 字符
	runes := []rune(str)

	// 检查索引范围
	if start < 0 {
		panic("Invalid start or end index")
	}

	// 提取子字符串
	return string(runes[start:len(runes)])
}
