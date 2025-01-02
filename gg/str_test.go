package gg

import (
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
