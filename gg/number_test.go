package gg

import (
	"fmt"
	"testing"
)

func TestNumber01(t *testing.T) {
	// 示例：判断值是否为数字类型
	var intValue int = 42
	var floatValue float64 = 3.14
	var stringValue string = "123"

	fmt.Printf("Is %v a number? %v\n", intValue, Number_isNumber(intValue))
	fmt.Printf("Is %v a number? %v\n", floatValue, Number_isNumber(floatValue))
	fmt.Printf("Is %v a number? %v\n", stringValue, Number_isNumberByString(stringValue))
}
