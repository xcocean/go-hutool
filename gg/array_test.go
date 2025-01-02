package gg

import (
	"fmt"
	"testing"
)

func TestArray01(t *testing.T) {
	// 示例：将新元素添加到已有切片中
	intBuffer := []int{1, 2, 3}
	newInts := []int{4, 5}
	resultInt := Array_append(intBuffer, newInts...)
	fmt.Printf("Original int buffer: %v\n", intBuffer)
	fmt.Printf("New int buffer: %v\n", resultInt)

	stringBuffer := []string{"Hello", "World"}
	newStrings := []string{"Go", "Lang"}
	resultString := Array_append(stringBuffer, newStrings...)
	fmt.Printf("Original string buffer: %v\n", stringBuffer)
	fmt.Printf("New string buffer: %v\n", resultString)

	type Person struct {
		Name string
		Age  int
	}
	personBuffer := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	newPersons := []Person{
		{"Charlie", 35},
	}
	resultPerson := Array_append(personBuffer, newPersons...)
	fmt.Printf("Original person buffer: %v\n", personBuffer)
	fmt.Printf("New person buffer: %v\n", resultPerson)
}

func TestArray02(t *testing.T) {
	// 示例：将新元素插入到已有切片中的指定位置
	intBuffer := []int{1, 2, 3}
	newInts := []int{4, 5}
	resultInt := Array_insert(intBuffer, 1, newInts...)
	fmt.Printf("Original int buffer: %v\n", intBuffer)
	fmt.Printf("New int buffer: %v\n", resultInt)

	stringBuffer := []string{"Hello", "World"}
	newStrings := []string{"Go", "Lang"}
	resultString := Array_insert(stringBuffer, 1, newStrings...)
	fmt.Printf("Original string buffer: %v\n", stringBuffer)
	fmt.Printf("New string buffer: %v\n", resultString)

	type Person struct {
		Name string
		Age  int
	}
	personBuffer := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	newPersons := []Person{
		{"Charlie", 35},
	}
	resultPerson := Array_insert(personBuffer, 1, newPersons...)
	fmt.Printf("Original person buffer: %v\n", personBuffer)
	fmt.Printf("New person buffer: %v\n", resultPerson)
}

func TestArray03(t *testing.T) {
	// 示例：生成一个新的重新设置大小的切片
	intData := []int{1, 2, 3}
	resizedIntData := Array_resize(intData, 5)
	fmt.Printf("Original int data: %v\n", intData)
	fmt.Printf("Resized int data: %v\n", resizedIntData)

	stringData := []string{"Hello", "World"}
	resizedStringData := Array_resize(stringData, 4)
	fmt.Printf("Original string data: %v\n", stringData)
	fmt.Printf("Resized string data: %v\n", resizedStringData)

	type Person struct {
		Name string
		Age  int
	}
	personData := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	resizedPersonData := Array_resize(personData, 3)
	fmt.Printf("Original person data: %v\n", personData)
	fmt.Printf("Resized person data: %v\n", resizedPersonData)
}

func TestArray04(t *testing.T) {
	// 示例：复制切片
	intSlice := []int{1, 2, 3}
	copiedIntSlice := Array_copy(intSlice)
	fmt.Printf("Original int slice: %v\n", intSlice)
	fmt.Printf("Copied int slice: %v\n", copiedIntSlice)

	stringSlice := []string{"Hello", "World"}
	copiedStringSlice := Array_copy(stringSlice)
	fmt.Printf("Original string slice: %v\n", stringSlice)
	fmt.Printf("Copied string slice: %v\n", copiedStringSlice)

	type Person struct {
		Name string
		Age  int
	}
	personSlice := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	copiedPersonSlice := Array_copy(personSlice)
	fmt.Printf("Original person slice: %v\n", personSlice)
	fmt.Printf("Copied person slice: %v\n", copiedPersonSlice)
}

func TestArray05(t *testing.T) {
	// 示例：将多个切片合并在一起
	intArray1 := []int{1, 2, 3}
	intArray2 := []int{4, 5}
	intArray3 := []int{6, 7, 8}
	mergedIntArray := Array_addAll(intArray1, intArray2, intArray3)
	fmt.Printf("Merged int array: %v\n", mergedIntArray)

	stringArray1 := []string{"Hello", "World"}
	stringArray2 := []string{"Go", "Lang"}
	mergedStringArray := Array_addAll(stringArray1, stringArray2)
	fmt.Printf("Merged string array: %v\n", mergedStringArray)

	type Person struct {
		Name string
		Age  int
	}
	personArray1 := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	personArray2 := []Person{
		{"Charlie", 35},
	}
	mergedPersonArray := Array_addAll(personArray1, personArray2)
	fmt.Printf("Merged person array: %v\n", mergedPersonArray)
}

func TestArray06(t *testing.T) {
	// 示例：去除切片中的 nil 元素
	var nilPtr *int
	intArray := []*int{nilPtr, new(int), new(int)}
	nonNullIntArray := Array_removeNull(intArray)
	fmt.Printf("Original int array: %v\n", intArray)
	fmt.Printf("Non-null int array: %v\n", nonNullIntArray)

	var nilString *string
	stringArray := []*string{nilString, new(string), new(string)}
	nonNullStringArray := Array_removeNull(stringArray)
	fmt.Printf("Original string array: %v\n", stringArray)
	fmt.Printf("Non-null string array: %v\n", nonNullStringArray)

	type Person struct {
		Name string
		Age  int
	}
	var nilPerson *Person
	personArray := []*Person{nilPerson, &Person{"Alice", 25}, &Person{"Bob", 30}}
	nonNullPersonArray := Array_removeNull(personArray)
	fmt.Printf("Original person array: %v\n", personArray)
	fmt.Printf("Non-null person array: %v\n", nonNullPersonArray)
}

func TestArray07(t *testing.T) {
	a := []string{"", "asd", "123"}
	fmt.Printf("array: %v\n", Array_removeNull(a))

	aa := []int{0, 1, 2, 3}
	fmt.Printf("array: %v\n", Array_removeNull(aa))

	aaa := []int64{0, 1, 2, 3}
	fmt.Printf("array: %v\n", Array_removeNull(aaa))
}

func TestArray08(t *testing.T) {
	// 示例：判断数组或切片是否包含某个元素
	intArray := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains 3 in intArray? %v\n", Array_contains(intArray, 3))
	fmt.Printf("Contains 6 in intArray? %v\n", Array_contains(intArray, 6))

	stringArray := []string{"Hello", "World"}
	fmt.Printf("Contains 'World' in stringArray? %v\n", Array_contains(stringArray, "World"))
	fmt.Printf("Contains 'Go' in stringArray? %v\n", Array_contains(stringArray, "Go"))

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	fmt.Printf("Contains {Bob 30} in personArray? %v\n", Array_contains(personArray, Person{"Bob", 30}))
	fmt.Printf("Contains {Charlie 35} in personArray? %v\n", Array_contains(personArray, Person{"Charlie", 35}))
}

func TestArray09(t *testing.T) {
	// 示例：获取数组或切片中指定索引的值
	intArray := []int{1, 2, 3, 4, 5}
	fmt.Printf("Value at index 2: %v\n", Array_get(intArray, 2))
	fmt.Printf("Value at index -1: %v\n", Array_get(intArray, -1))
	fmt.Printf("Value at index 10: %v\n", Array_get(intArray, 10))

	stringArray := []string{"Hello", "World", "Go"}
	fmt.Printf("Value at index 1: %v\n", Array_get(stringArray, 1))
	fmt.Printf("Value at index -2: %v\n", Array_get(stringArray, -2))
	fmt.Printf("Value at index -5: %v\n", Array_get(stringArray, -5))

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	fmt.Printf("Value at index 0: %v\n", Array_get(personArray, 0))
	fmt.Printf("Value at index -3: %v\n", Array_get(personArray, -3))
	fmt.Printf("Value at index 3: %v\n", Array_get(personArray, 3))
}

func TestArray10(t *testing.T) {
	// 示例：获取子数组（子切片）
	intArray := []int{1, 2, 3, 4, 5}
	subIntArray := Array_subArray(intArray, 1, 4)
	fmt.Printf("Original int array: %v\n", intArray)
	fmt.Printf("Sub int array (1, 4): %v\n", subIntArray)

	stringArray := []string{"Hello", "World", "Go", "Lang"}
	subStringArray := Array_subArray(stringArray, -3, -1)
	fmt.Printf("Original string array: %v\n", stringArray)
	fmt.Printf("Sub string array (-3, -1): %v\n", subStringArray)

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	subPersonArray := Array_subArray(personArray, 0, 2)
	fmt.Printf("Original person array: %v\n", personArray)
	fmt.Printf("Sub person array (0, 2): %v\n", subPersonArray)
}

func TestArray11(t *testing.T) {
	// 示例：移除数组中指定的元素
	intArray := []int{1, 2, 3, 2, 4}
	removedIntArray := Array_removeAll(intArray, 2)
	fmt.Printf("Original int array: %v\n", intArray)
	fmt.Printf("After removing 2: %v\n", removedIntArray)

	stringArray := []string{"Hello", "World", "Go", "World"}
	removedStringArray := Array_removeEle(stringArray, "World")
	fmt.Printf("Original string array: %v\n", stringArray)
	fmt.Printf("After removing 'World': %v\n", removedStringArray)

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Alice", 25},
	}
	removedPersonArray := Array_removeAll(personArray, Person{"Alice", 25})
	fmt.Printf("Original person array: %v\n", personArray)
	fmt.Printf("After removing {Alice 25}: %v\n", removedPersonArray)
}

func TestArray12(t *testing.T) {
	// 示例：移除数组中所有匹配的元素
	intArray := []int{1, 2, 3, 2, 4}
	removedIntArray := Array_removeAll(intArray, 2)
	fmt.Printf("Original int array: %v\n", intArray)
	fmt.Printf("After removing all 2: %v\n", removedIntArray)

	stringArray := []string{"Hello", "World", "Go", "World"}
	removedStringArray := Array_removeAll(stringArray, "World")
	fmt.Printf("Original string array: %v\n", stringArray)
	fmt.Printf("After removing all 'World': %v\n", removedStringArray)

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Alice", 25},
	}
	removedPersonArray := Array_removeAll(personArray, Person{"Alice", 25})
	fmt.Printf("Original person array: %v\n", personArray)
	fmt.Printf("After removing all {Alice 25}: %v\n", removedPersonArray)
}

func TestArray13(t *testing.T) {
	// 示例：交换数组中两个位置的值
	intArray := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original int array: %v\n", intArray)
	Array_swap(intArray, 1, 3)
	fmt.Printf("After swapping index 1 and 3: %v\n", intArray)

	stringArray := []string{"Hello", "World", "Go", "Lang"}
	fmt.Printf("Original string array: %v\n", stringArray)
	Array_swap(stringArray, 0, 2)
	fmt.Printf("After swapping index 0 and 2: %v\n", stringArray)

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	fmt.Printf("Original person array: %v\n", personArray)
	Array_swap(personArray, 1, 2)
	fmt.Printf("After swapping index 1 and 2: %v\n", personArray)
}

func TestArray14(t *testing.T) {
	// 示例：去重数组中的元素
	intArray := []int{1, 2, 2, 3, 4, 4, 5}
	distinctIntArray := Array_distinct(intArray)
	fmt.Printf("Original int array: %v\n", intArray)
	fmt.Printf("Distinct int array: %v\n", distinctIntArray)

	stringArray := []string{"Hello", "World", "Hello", "Go"}
	distinctStringArray := Array_distinct(stringArray)
	fmt.Printf("Original string array: %v\n", stringArray)
	fmt.Printf("Distinct string array: %v\n", distinctStringArray)

	type Person struct {
		Name string
		Age  int
	}
	personArray := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Alice", 25},
	}
	distinctPersonArray := Array_distinct(personArray)
	fmt.Printf("Original person array: %v\n", personArray)
	fmt.Printf("Distinct person array: %v\n", distinctPersonArray)
}
