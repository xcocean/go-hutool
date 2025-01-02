package gg

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

// 生成一个随机的布尔值
func random_randomBoolean() bool {
	// 生成一个随机字节
	var b [1]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random bool: %s", err))
	}
	// 判断字节的最低有效位是 0 还是 1
	return b[0]&1 == 1
}

// RandomInt 生成一个随机的 int 值
func Random_randomInt(min, max int) int {
	if min >= max {
		panic("Invalid range: min must be less than max")
	}
	// 生成一个范围在 [0, max-min) 的随机数
	rangeSize := big.NewInt(int64(max - min + 1))
	randomNum, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random number: %s", err))
	}
	return min + int(randomNum.Int64())
}

// RandomBytes 生成指定长度的随机字节切片
func Random_randomBytes(length int) []byte {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random bytes: %s", err))
	}
	return bytes
}

// RandomFloat 生成一个随机的 float32 值
func Random_randomFloat() float32 {
	var buf [4]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random float32: %s", err))
	}
	// 将字节转换为 uint32，再转换为 float32
	bits := binary.BigEndian.Uint32(buf[:])
	return math.Float32frombits(bits)
}

// RandomDouble 生成一个随机的 float64 值
func Random_randomDouble() float64 {
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random float64: %s", err))
	}
	// 将字节转换为 uint64，再转换为 float64
	bits := binary.BigEndian.Uint64(buf[:])
	return math.Float64frombits(bits)
}

// 生成随机字符串
func Random_randomString(length int) string {
	bytes := Random_randomBytes(length)
	for i, b := range bytes {
		bytes[i] = DEFAULT_ALPHABET[b%byte(len(DEFAULT_ALPHABET))]
	}
	return string(bytes)
}

// 生成一个加密安全的随机 uint32 值
func Random_uint32() uint32 {
	var b [4]byte // uint32 需要 4 个字节
	_, err := rand.Read(b[:])
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random uint32: %s", err))
	}
	// 将字节转换为 uint32
	return binary.BigEndian.Uint32(b[:])
}

// 生成指定范围的随机 int 值，参数无效时 panic
func Random_randomInt2(min, max int) int {
	if min >= max {
		panic("Invalid range: min must be less than max")
	}
	// 生成一个范围在 [0, max-min) 的随机数
	rangeSize := big.NewInt(int64(max - min + 1))
	randomNum, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random number: %s", err))
	}
	return min + int(randomNum.Int64())
}

// 生成指定范围的随机 int64 值，参数无效时 panic
func Random_randomLong2(min, max int64) int64 {
	if min >= max {
		panic("Invalid range: min must be less than max")
	}
	// 生成一个范围在 [0, max-min) 的随机数
	rangeSize := big.NewInt(max - min + 1)
	randomNum, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random number: %s", err))
	}
	return min + randomNum.Int64()
}

// 生成指定长度的随机字节切片，参数无效时 panic
func Random_randomBytes2(length int) []byte {
	if length <= 0 {
		panic("Invalid length: length must be greater than 0")
	}
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random bytes: %s", err))
	}
	return bytes
}

// 生成指定范围的随机 float32 值，参数无效时 panic
func Random_randomFloat2(min, max float32) float32 {
	if min >= max {
		panic("Invalid range: min must be less than max")
	}
	// 生成一个随机的 float32 值
	var buf [4]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random float32: %s", err))
	}
	// 将字节转换为 uint32，再转换为 float32
	bits := binary.BigEndian.Uint32(buf[:])
	value := math.Float32frombits(bits)
	// 将值映射到指定范围
	return min + (max-min)*value
}

// 生成指定范围的随机 float64 值，参数无效时 panic
func Random_randomDouble2(min, max float64) float64 {
	if min >= max {
		panic("Invalid range: min must be less than max")
	}
	// 生成一个随机的 float64 值
	var buf [8]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random float64: %s", err))
	}
	// 将字节转换为 uint64，再转换为 float64
	bits := binary.BigEndian.Uint64(buf[:])
	value := math.Float64frombits(bits)
	// 将值映射到指定范围
	return min + (max-min)*value
}

// 从指定字符串中随机获取指定长度的字符串
func Random_stringFromSource(source string, length int) string {
	if length <= 0 {
		panic("Invalid length: length must be greater than 0")
	}
	if len(source) == 0 {
		panic("Invalid source: source string must not be empty")
	}

	// 将字符串转换为 rune 切片以便处理 Unicode 字符
	sourceRunes := []rune(source)
	result := make([]rune, length)

	for i := 0; i < length; i++ {
		// 生成一个随机索引
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(sourceRunes))))
		if err != nil {
			panic(fmt.Sprintf("Failed to generate random index: %s", err))
		}
		// 从源字符串中获取随机字符
		result[i] = sourceRunes[randomIndex.Int64()]
	}

	return string(result)
}

// 从指定数组中随机获取指定长度的子数组
func Random_randomSubArray[T any](source []T, length int) []T {
	if length <= 0 {
		panic("Invalid length: length must be greater than 0")
	}
	if len(source) == 0 {
		panic("Invalid source: source array must not be empty")
	}
	if length > len(source) {
		panic("Invalid length: length must be less than or equal to the source array length")
	}

	// 创建一个副本以避免修改原数组
	sourceCopy := make([]T, len(source))
	copy(sourceCopy, source)

	result := make([]T, length)
	for i := 0; i < length; i++ {
		// 生成一个随机索引
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(sourceCopy))))
		if err != nil {
			panic(fmt.Sprintf("Failed to generate random index: %s", err))
		}
		// 将选中的元素添加到结果中
		result[i] = sourceCopy[randomIndex.Int64()]
		// 从源数组中移除已选中的元素，避免重复
		sourceCopy = append(sourceCopy[:randomIndex.Int64()], sourceCopy[randomIndex.Int64()+1:]...)
	}

	return result
}

// 从指定数组中随机获取一个元素
func Random_randomElement[T any](source []T) T {
	if len(source) == 0 {
		panic("Invalid source: source array must not be empty")
	}

	// 生成一个随机索引
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(source))))
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random index: %s", err))
	}

	// 返回随机元素
	return source[randomIndex.Int64()]
}
