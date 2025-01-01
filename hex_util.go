package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// IsHexNumber 判断字符串是否是有效的十六进制数
// '0x123abc' --> true
// '0XABCDEF' --> true
// '123abc' --> true
// 'ABCDEF' --> true
// '0x123g' --> false
// '123g' --> false
// ” --> false
// '0x' --> false
func IsHexNumber(s string) bool {
	// 去除前缀 "0x" 或 "0X"
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		s = s[2:]
	}

	// 空字符串不是有效的十六进制数
	if len(s) == 0 {
		return false
	}

	// 检查每个字符是否是有效的十六进制字符
	for _, r := range s {
		if !unicode.IsDigit(r) && !(r >= 'a' && r <= 'f') && !(r >= 'A' && r <= 'F') {
			return false
		}
	}
	return true
}

// BytesToHexString 将字节数组转换为十六进制字符串
func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

// HexArrayToString 将十六进制字符数组转换为字符串
// [48 65 6C 6C 6F 20 57 6F 72 6C 64]  --> Hello World
func HexArrayToString(hexArray []string) (string, error) {
	// 拼接字符数组为一个完整的十六进制字符串
	hexString := strings.Join(hexArray, "")

	// 解码十六进制字符串为字节数组
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	// 将字节数组转换为字符串
	return string(bytes), nil
}

// HexArrayToBytes 将十六进制字符数组转换为字节数组
func HexArrayToBytes(hexArray []string) ([]byte, error) {
	// 拼接字符数组为一个完整的十六进制字符串
	hexString := strings.Join(hexArray, "")

	// 解码十六进制字符串为字节数组
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func IntToHexString(n int) string {
	return fmt.Sprintf("%x", n) // %x 表示小写十六进制，%X 表示大写十六进制
}

func HexStringToInt(hexString string) (int, error) {
	// 16 表示十六进制，64 表示返回 int64
	value, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(value), nil
}
