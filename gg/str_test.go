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
	println(IsBlank(str) == true)
	println(IsNotBlank(str) == false)
	println(IsBlank(" \t\n") == true)

	str = "  "
	println(IsBlank(str) == true)
	println(IsNotBlank(str) == false)

}

func TestEmpty(t *testing.T) {
	var str string
	println(IsEmpty(str) == true)
	println(IsNotEmpty(str) == false)
	println(IsEmpty(" \t\n") == true)

	str = "  "
	println(IsEmpty(str) == false)
	println(IsNotEmpty(str) == false)

	str = "  "
	println(IsEmpty(str) == false)
	println(IsNotEmpty(str) == false)
}

func TestTrim(t *testing.T) {
	assert(Trim("") == "", "Trim")
	assert(Trim("     ") == "", "Trim")
	assert(Trim("abc") == "abc", "Trim")
	assert(Trim("    abc    ") == "abc", "Trim")

	assert(TrimStart("    abc    ") == "abc    ", "TrimStart")
	assert(TrimStart("abc    ") == "abc    ", "TrimStart")
	assert(TrimEnd("abc    ") == "abc", "TrimEnd")
}

func TestWith(t *testing.T) {
	assert(StartWith("asd", "as") == true, "StartWith")
	assert(StartWith("asd", "1as") == false, "StartWith")
	assert(StartWith("", "123") == false, "StartWith")

	assert(EndWith("asd", "sd") == true, "EndWith")
	assert(EndWith("asd", "aasd") == false, "EndWith")
	assert(EndWith("", "") == false, "EndWith")
}

func TestToString(t *testing.T) {
	var arr = [5]int{10, 20, 30, 40, 50}
	println(ToString(arr))

	arr2 := [2]int{10, 20}
	println(ToString(arr2))
}
