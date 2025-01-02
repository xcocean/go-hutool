package gg

import (
	"fmt"
	"testing"
)

func TestRandom01(t *testing.T) {
	// 示例：生成随机值
	fmt.Println("Random Int:", Random_randomInt(1, 100))
	fmt.Println("Random Bytes:", Random_randomBytes(8))
	fmt.Println("Random Float32:", Random_randomFloat())
	fmt.Println("Random Float64:", Random_randomDouble())
	fmt.Println("Random_randomString:", Random_randomString(10))
	fmt.Println("random_randomBoolean:", random_randomBoolean())
}

func TestRandom02(t *testing.T) {
	// 示例：生成随机值
	fmt.Println("Random Int:", Random_randomInt(1, 100))
	fmt.Println("Random Long:", Random_randomLong2(1000, 9999))
	fmt.Println("Random Bytes:", Random_randomBytes2(8))
	fmt.Println("Random Float32:", Random_randomFloat2(0.0, 1.0))
	fmt.Println("Random Float64:", Random_randomDouble2(0.0, 1.0))
}

func TestRandom03(t *testing.T) {
	// 示例：从指定字符串中随机获取指定长度的字符串
	source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 10
	result := Random_stringFromSource(source, length)
	fmt.Printf("Random String: %s\n", result)
}

func TestRandom04(t *testing.T) {
	// 示例：从指定数组中随机获取指定长度的子数组
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	length := 5
	result := Random_randomSubArray(source, length)
	fmt.Printf("Random Sub Array: %v\n", Str_toString(result))
}

func TestRandom05(t *testing.T) {
	source := []string{"apple", "banana", "cherry", "date", "elderberry"}
	length := 3
	result := Random_randomSubArray(source, length)
	fmt.Printf("Random Sub Array: %v\n", result)
}

func TestRandom06(t *testing.T) {
	// 示例：从整数数组中随机获取一个元素
	intArray := []int{1, 2, 3, 4, 5}
	intElement := Random_randomElement(intArray)
	fmt.Printf("Random Int Element: %d\n", intElement)

	// 示例：从字符串数组中随机获取一个元素
	stringArray := []string{"apple", "banana", "cherry", "date", "elderberry"}
	stringElement := Random_randomElement(stringArray)
	fmt.Printf("Random String Element: %s\n", stringElement)

	// 示例：从结构体数组中随机获取一个元素
	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	personElement := Random_randomElement(personArray)
	fmt.Printf("Random Person Element: %+v\n", personElement)
}
