package gg

import (
	"fmt"
	"testing"
)

func TestChar01(t *testing.T) {
	// 示例：判断字符是否为 ASCII 字符
	c1 := 'A'
	c2 := 'é'
	fmt.Printf("Is '%c' ASCII? %v\n", c1, Char_isAscii(c1))
	fmt.Printf("Is '%c' ASCII? %v\n", c2, Char_isAscii(c2))

	// 示例：判断字符是否为可见 ASCII 字符
	ch1 := 'A'
	ch2 := '\t' // 制表符（不可见字符）
	ch3 := '~'
	ch4 := 'é' // 非 ASCII 字符
	fmt.Printf("Is '%c' printable ASCII? %v\n", ch1, Char_isAsciiPrintable(ch1))
	fmt.Printf("Is '%c' printable ASCII? %v\n", ch2, Char_isAsciiPrintable(ch2))
	fmt.Printf("Is '%c' printable ASCII? %v\n", ch3, Char_isAsciiPrintable(ch3))
	fmt.Printf("Is '%c' printable ASCII? %v\n", ch4, Char_isAsciiPrintable(ch4))
}

func TestChar02(t *testing.T) {
	// 示例：判断字符是否为小写字母
	ch1 := 'a'
	ch2 := 'A'
	ch3 := '1'
	ch4 := 'é' // 小写字母（带重音符号）
	fmt.Printf("Is '%c' a lowercase letter? %v\n", ch1, Char_isLetterLower(ch1))
	fmt.Printf("Is '%c' a lowercase letter? %v\n", ch2, Char_isLetterLower(ch2))
	fmt.Printf("Is '%c' a lowercase letter? %v\n", ch3, Char_isLetterLower(ch3))
	fmt.Printf("Is '%c' a lowercase letter? %v\n", ch4, Char_isLetterLower(ch4))
}

func TestChar03(t *testing.T) {
	// 示例：判断字符是否为大写字母
	ch1 := 'A'
	ch2 := 'a'
	ch3 := '1'
	ch4 := 'É' // 大写字母（带重音符号）
	fmt.Printf("Is '%c' an uppercase letter? %v\n", ch1, Char_isLetterUpper(ch1))
	fmt.Printf("Is '%c' an uppercase letter? %v\n", ch2, Char_isLetterUpper(ch2))
	fmt.Printf("Is '%c' an uppercase letter? %v\n", ch3, Char_isLetterUpper(ch3))
	fmt.Printf("Is '%c' an uppercase letter? %v\n", ch4, Char_isLetterUpper(ch4))
}
