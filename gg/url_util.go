package gg

import (
	"fmt"
	"net/url"
)

// 对字符串进行 URL 编码
func URL_Encode(input string) string {
	return url.QueryEscape(input)
}

// 对 URL 编码的字符串进行解码
func URL_Decode(encoded string) string {
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		panic(fmt.Sprintf("URL decode failed: %s", err))
	}
	return decoded
}
