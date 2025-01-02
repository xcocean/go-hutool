package gg

import "reflect"

// 判断数组或切片是否为空
func Array_isEmpty[T any](array []T) bool {
	return len(array) == 0
}

// 如果给定数组为空，返回默认数组
func Array_defaultIfEmpty[T any](array, defaultArray []T) []T {
	if len(array) == 0 {
		return defaultArray
	}
	return array
}

func Array_isNotEmpty[T any](array []T) bool {
	return !Array_isEmpty(array)
}

// 是否包含null元素
func Array_hasNull[T any](array []T) bool {
	for _, item := range array {
		if reflect.ValueOf(item).IsNil() {
			return true
		}
	}
	return false
}

// 将新元素添加到已有切片中，返回一个新的切片
func Array_append[T any](buffer []T, newElements ...T) []T {
	return append(buffer, newElements...)
}

// 将新元素插入到已有切片中的指定位置，返回一个新的切片
func Array_insert[T any](buffer []T, index int, newElements ...T) []T {
	// 检查索引是否有效
	if index < 0 || index > len(buffer) {
		panic("Index out of range")
	}

	// 创建一个新的切片，长度为原切片长度加上新元素的长度
	newBuffer := make([]T, len(buffer)+len(newElements))

	// 将原切片的前半部分复制到新切片中
	copy(newBuffer, buffer[:index])

	// 将新元素插入到新切片中
	copy(newBuffer[index:], newElements)

	// 将原切片的后半部分复制到新切片中
	copy(newBuffer[index+len(newElements):], buffer[index:])

	return newBuffer
}

// 生成一个新的重新设置大小的切片
func Array_resize[T any](data []T, newSize int) []T {
	// 创建一个新的切片
	newData := make([]T, newSize)

	// 复制原切片的元素到新切片中
	copy(newData, data)

	return newData
}

// 复制切片，返回一个新的切片
func Array_copy[T any](src []T) []T {
	// 创建一个新的切片，长度与源切片相同
	dst := make([]T, len(src))

	// 复制源切片的元素到新切片中
	copy(dst, src)

	return dst
}

// 将多个切片合并在一起，返回一个新的切片
func Array_addAll[T any](arrays ...[]T) []T {
	// 计算所有切片的长度总和
	totalLen := 0
	for _, array := range arrays {
		totalLen += len(array)
	}

	// 创建一个新的切片，长度为所有切片的长度总和
	result := make([]T, 0, totalLen)

	// 将所有切片的元素合并到新切片中
	for _, array := range arrays {
		result = append(result, array...)
	}

	return result
}

// 去除切片中的 nil 元素，返回一个新的切片
func Array_removeNull[T any](array []T) []T {
	// 创建一个新的切片，用于存储非 nil 元素
	result := make([]T, 0, len(array))

	// 遍历原切片，过滤掉 nil 元素
	for _, item := range array {
		if !isNil(item) {
			result = append(result, item)
		}
	}

	return result
}

// 检查一个值是否为 nil
func isNil[T any](value T) bool {
	// 使用反射检查值是否为 nil
	return any(value) == nil || (isPointer(value) && reflect.ValueOf(value).IsNil())
}

// 检查一个值是否为指针类型
func isPointer[T any](value T) bool {
	return reflect.TypeOf(value).Kind() == reflect.Ptr
}

// 判断数组或切片是否包含某个元素
func Array_contains[T any](array []T, value T) bool {
	for _, item := range array {
		if reflect.DeepEqual(item, value) {
			return true
		}
	}
	return false
}

// 获取数组或切片中指定索引的值，支持负数索引，例如-1表示倒数第一个值，越界返回 nil
func Array_get[T any](array []T, index int) interface{} {
	// 处理负数索引
	if index < 0 {
		index += len(array)
	}

	// 检查索引是否越界
	if index < 0 || index >= len(array) {
		return nil
	}

	// 返回指定索引的值
	return array[index]
}

// Sub 获取子数组（子切片），支持负数索引
func Array_subArray[T any](array []T, start, end int) []T {
	// 处理负数索引
	if start < 0 {
		start += len(array)
	}
	if end < 0 {
		end += len(array)
	}

	// 检查索引范围
	if start < 0 || end > len(array) || start > end {
		panic("Invalid start or end index")
	}

	// 返回子切片
	return array[start:end]
}

// RemoveEle 移除数组中指定的元素，只移除第一个匹配的元素
func Array_removeEle[T any](array []T, element T) []T {
	// 创建一个新的切片，用于存储移除后的元素
	result := make([]T, 0, len(array))

	// 标记是否已经移除了第一个匹配的元素
	removed := false

	// 遍历原切片
	for _, item := range array {
		// 如果未移除且当前元素与目标元素相等，跳过该元素
		if !removed && reflect.DeepEqual(item, element) {
			removed = true
			continue
		}
		// 将其他元素添加到新切片中
		result = append(result, item)
	}

	return result
}

// 移除数组中所有匹配的元素
func Array_removeAll[T any](array []T, element T) []T {
	// 创建一个新的切片，用于存储移除后的元素
	result := make([]T, 0, len(array))

	// 遍历原切片，过滤掉所有匹配的元素
	for _, item := range array {
		if !reflect.DeepEqual(item, element) {
			result = append(result, item)
		}
	}

	return result
}

// 交换数组中两个位置的值
func Array_swap[T any](array []T, index1, index2 int) []T {
	// 检查索引是否有效
	if index1 < 0 || index1 >= len(array) || index2 < 0 || index2 >= len(array) {
		panic("Index out of range")
	}

	// 交换两个位置的值
	array[index1], array[index2] = array[index2], array[index1]

	return array
}

// 去重数组中的元素，返回一个新的数组
func Array_distinct[T comparable](array []T) []T {
	// 使用 map 来记录已经出现的元素
	seen := make(map[T]bool)
	result := make([]T, 0, len(array))

	// 遍历原切片，过滤掉重复元素
	for _, item := range array {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
