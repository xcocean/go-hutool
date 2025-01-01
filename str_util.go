package main

import (
	"fmt"
	"strings"
)

// 字符串是否为空白
// IsBlank("") // true
// IsBlank(" \t\n") // true
// IsBlank("abc") // false
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 字符串是否为非空白
// IsNotBlank("") // false
// IsNotBlank("") // false
// IsNotBlank(" \t\n") // false
// IsNotBlank("abc") // true
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// 字符串是否为空
// IsEmpty("") // true
// IsEmpty(" \t\n") // false
// IsEmpty("abc") // false
func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// 字符串是否为非空白
// IsNotEmpty("") // false
// IsNotEmpty(" \t\n") // true
// IsNotEmpty("abc") // true
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// 是否包含空字符串
// 如果指定的字符串数组的长度为 0，或者其中的任意一个元素是空字符串，则返回 true。
func hasEmpty(str ...string) bool {
	if len(str) == 0 {
		return true
	}
	for _, s := range str {
		if IsEmpty(s) {
			return true
		}
	}
	return false
}

// 指定字符串数组中的元素，是否全部为空字符串。
func IsAllEmpty(str ...string) bool {
	if len(str) == 0 {
		return true
	}
	for _, s := range str {
		if !IsEmpty(s) {
			return false
		}
	}
	return true
}

// 指定字符串数组中的元素，是否全部为空字符串。
func IsAllBlank(str ...string) bool {
	if len(str) == 0 {
		return true
	}
	for _, s := range str {
		if !IsBlank(s) {
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
func Trim(str string) string {
	return strings.TrimSpace(str)
}

/**
 * 除去字符串头部的空白，如果字符串是{@code ""}，则返回{@code ""}。
 */
func TrimStart(str string) string {
	return strings.TrimLeft(str, " ")
}

/**
 * 除去字符串尾部的空白，如果字符串是{@code ""}，则返回{@code ""}。
 */
func TrimEnd(str string) string {
	return strings.TrimRight(str, " ")
}

// 是否以指定字符串开头
func StartWith(str string, c string) bool {
	if IsEmpty(str) {
		return false
	}
	return strings.HasPrefix(str, c)
}

// 字符串是否以给定字符结尾
func EndWith(str string, c string) bool {
	if IsEmpty(str) {
		return false
	}
	return strings.HasSuffix(str, c)
}

// 是否以指定字符串开头，忽略大小写
func StartWithIgnoreCase(s, prefix string) bool {
	if IsEmpty(s) {
		return false
	}
	return StartWith(strings.ToLower(s), strings.ToLower(prefix))
}

// 是否以指定字符串结尾，忽略大小写
func EndWithIgnoreCase(s, suffix string) bool {
	if IsEmpty(s) {
		return false
	}
	return EndWith(strings.ToLower(s), strings.ToLower(suffix))
}

func Equals(s, c string) bool {
	return s == c
}

func EqualsIgnoreCase(s, c string) bool {
	return strings.ToLower(s) == strings.ToLower(c)
}

// 字节数组转UTF-8字符串
func Utf8Str(byt []byte) string {
	return string(byt)
}

func ToString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}
