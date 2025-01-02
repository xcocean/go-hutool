package gg

import (
	"fmt"
	"testing"
)

func assert(bol bool, msg string) {
	if !bol {
		panic(msg)
	}
}

func TestBlank(t *testing.T) {
	var str string
	println(Str_isBlank(str) == true)
	println(Str_isNotBlank(str) == false)
	println(Str_isBlank(" \t\n") == true)

	str = "  "
	println(Str_isBlank(str) == true)
	println(Str_isNotBlank(str) == false)

}

func TestEmpty(t *testing.T) {
	var str string
	println(Str_isEmpty(str) == true)
	println(Str_isNotEmpty(str) == false)
	println(Str_isEmpty(" \t\n") == true)

	str = "  "
	println(Str_isEmpty(str) == false)
	println(Str_isNotEmpty(str) == false)

	str = "  "
	println(Str_isEmpty(str) == false)
	println(Str_isNotEmpty(str) == false)
}

func TestStr_trim(t *testing.T) {
	assert(Str_Trim("") == "", "Str_trim")
	assert(Str_Trim("     ") == "", "Str_trim")
	assert(Str_Trim("abc") == "abc", "Str_trim")
	assert(Str_Trim("    abc    ") == "abc", "Str_trim")

	assert(Str_trimStart("    abc    ") == "abc    ", "Str_trimStart")
	assert(Str_trimStart("abc    ") == "abc    ", "Str_trimStart")
	assert(Str_trimEnd("abc    ") == "abc", "Str_trimEnd")
}

func TestWith(t *testing.T) {
	assert(Str_startWith("asd", "as") == true, "Str_startWith")
	assert(Str_startWith("asd", "1as") == false, "Str_startWith")
	assert(Str_startWith("", "123") == false, "Str_startWith")

	assert(Str_endWith("asd", "sd") == true, "Str_endWith")
	assert(Str_endWith("asd", "aasd") == false, "Str_endWith")
	assert(Str_endWith("", "") == false, "Str_endWith")
}

func TestToString(t *testing.T) {
	var arr = [5]int{10, 20, 30, 40, 50}
	println(Str_toString(arr))

	arr2 := [2]int{10, 20}
	println(Str_toString(arr2))
}

func TestStr_removeSpace(t *testing.T) {
	// 示例：删除字符串中的空白字符
	input := "Hello, \nWorld!\t This is a test."
	output := Str_removeSpace(input)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("After removal: %q\n", output)
}

func TestStr_removeSpace2(t *testing.T) {
	// 示例：替换字符
	input := "Hello, World!"
	oldChar := 'o'
	newChar := 'x'
	output := Str_replaceChar(input, oldChar, newChar)
	fmt.Printf("Original: %s\n", input)
	fmt.Printf("After replacing '%c' with '%c': %s\n", oldChar, newChar, output)

	// 示例：替换子字符串
	input = "Hello, World!"
	oldStr := "World"
	newStr := "Go"
	output = Str_replaceString(input, oldStr, newStr)
	fmt.Printf("Original: %s\n", input)
	fmt.Printf("After replacing '%s' with '%s': %s\n", oldStr, newStr, output)
}
